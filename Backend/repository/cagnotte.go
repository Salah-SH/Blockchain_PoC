package repository

import (
	"errors"
	"strconv"

	"github.com/cagnotteApp/Backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) UpdatePaiement(accAddr string, paiementoken string, status bool) error {
	success := ""
	if status {
		success = "true"
	} else {
		success = "expired"
	}
	client := r.client
	filter := bson.M{"accaddr": accAddr, "paiements.Paiementtoken": paiementoken}
	update := bson.M{"$set": bson.M{"paiements.$.status": success}}
	_ = client.Database(r.dbName).Collection("Users").FindOneAndUpdate(r.context, filter, update)
	return nil

}

func (r *MongoRepository) RegisterPaiement(accAddr string, paiementoken string, cagnotteName string, amount string) error {
	client := r.client

	cur, err := client.Database(r.dbName).Collection("Users").Find(r.context, bson.M{"accaddr": accAddr})
	returnedValues := []domain.User{}
	err = cur.All(r.context, &returnedValues)
	paiements := returnedValues[0].Paiements
	paiements = append(paiements, domain.Paiement{
		Paiementtoken: paiementoken,
		Status:        "not paid",
		Amount:        amount,
		CagnotteName:  cagnotteName,
	})
	_, err = client.Database(r.dbName).Collection("Users").UpdateOne(
		r.context,
		bson.M{"accaddr": accAddr},
		bson.D{
			{"$set", bson.D{{"paiements", paiements}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) RegisterAdminCreds(admin *domain.AdminCreds) (*domain.AdminCreds, error) {
	client := r.client
	_, err := client.Database(r.dbName).Collection("Admin").InsertOne(r.context, admin)

	if err != nil {
		return nil, err
	}
	return admin, nil

}

func (r *MongoRepository) GetAdminCreds() (domain.AdminCreds, error) {
	client := r.client
	adminCreds := domain.AdminCreds{}
	err := client.Database(r.dbName).Collection("Admin").FindOne(r.context, bson.M{}).Decode(&adminCreds)
	if err != nil {
		return adminCreds, err
	}
	return adminCreds, nil

}
func (r *MongoRepository) SaveCagnotte(cagnotte *domain.Cagnotte) (*domain.Cagnotte, error) {
	client := r.client
	result, err := client.Database(r.dbName).Collection("Cagnottes").InsertOne(r.context, cagnotte)
	if err != nil {
		return nil, err
	}
	cagnotte.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return cagnotte, nil
}

func (r *MongoRepository) FindCagnotte(name string) (*domain.Cagnotte, error) {

	client := r.client
	cur, err := client.Database(r.dbName).Collection("Cagnottes").Find(r.context, bson.M{"owner": name})
	if err != nil {
		return nil, err
	}
	returnedValues := []domain.Cagnotte{}

	err = cur.All(r.context, &returnedValues)

	if err != nil {
		return nil, err
	}
	if len(returnedValues) == 0 {
		return nil, errors.New("cagnotte not found in db")
	}
	return &returnedValues[0], nil
}

func (r *MongoRepository) UpdateCagnotte(cagnotte *domain.Cagnotte) (bool, error) {

	client := r.client

	result, err := client.Database(r.dbName).Collection("Cagnottes").UpdateOne(
		r.context,
		bson.M{"name": cagnotte.Name},
		bson.D{
			{"$set", bson.D{{"historic", cagnotte.Participation}}},
		},
	)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount != 0, nil
}

func (r *MongoRepository) SaveAddTx(txs *domain.TransactionDetails, collectionName string) (bool, error) {
	client := r.client

	returnedValues := []domain.TransactionDetails{}
	cur, err := client.Database(r.dbName).Collection(collectionName).Find(r.context, bson.M{"_id": txs.Hash})

	cur.All(r.context, &returnedValues)
	if len(returnedValues) > 0 {

		return true, nil
	}

	_, err = client.Database(r.dbName).Collection(collectionName).InsertOne(r.context, txs)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *MongoRepository) FindTx(collectionName string, status ...string) ([]domain.TransactionDetails, error) {
	client := r.client
	returnedValues := []domain.TransactionDetails{}
	if len(status) == 1 {
		cur, err := client.Database(r.dbName).Collection(collectionName).Find(r.context, bson.M{"status": status[0], "considered": false})
		if err != nil {
			return nil, err
		}
		err = cur.All(r.context, &returnedValues)
		if err != nil {
			return nil, err
		}

		if len(returnedValues) > 0 {
			for _, tx := range returnedValues {
				_, err := client.Database(r.dbName).Collection(collectionName).UpdateOne(
					r.context,
					bson.M{"_id": tx.Hash},
					bson.D{
						{"$set", bson.D{{"considered", true}}},
					},
				)
				if err != nil {
					return nil, err
				}

			}
		}
		return returnedValues, nil
	}

	pipeline := bson.D{
		{"executed", false},
		{"$or", []interface{}{
			bson.D{{"status", status[0]}},
			bson.D{{"status", status[1]}},
		}},
	}
	cur, err := client.Database(r.dbName).Collection(collectionName).Find(r.context, pipeline)
	if err != nil {
		return nil, err
	}
	err = cur.All(r.context, &returnedValues)
	if err != nil {
		return nil, err
	}

	return returnedValues, nil

}

func (r *MongoRepository) UpdateDB(txs domain.TransactionDetails, success bool, collectionName string) (bool, error) {
	client := r.client

	result, err := client.Database(r.dbName).Collection(collectionName).UpdateOne(
		r.context,
		bson.M{"_id": txs.Hash},
		bson.D{
			{"$set", bson.D{{"status", strconv.FormatBool(success)}}},
		},
	)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount != 0, nil

}

func (r *MongoRepository) UpdateValidationDB(txs domain.TransactionDetails, success bool, collectionName string, validationTxHash string) (bool, error) {
	client := r.client
	result, err := client.Database(r.dbName).Collection(collectionName).UpdateOne(
		r.context,
		bson.M{"_id": txs.Hash},
		bson.D{
			{"$set", bson.D{{"executed", success}, {"confimationonledger", false}, {"confirmtxhash", validationTxHash}}},
		},
	)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount != 0, nil
}

func (r *MongoRepository) GetConfig() (domain.Config, error) {
	client := r.client
	config := domain.Config{}
	names, err := client.Database(r.dbName).ListCollectionNames(r.context, bson.D{})
	initialise := true
	for _, name := range names {
		if name == "Config" {
			initialise = false
			break
		}
	}
	if initialise {
		config = domain.Config{HeightTx: 0, HeightValid: 0}
		_, err := client.Database(r.dbName).Collection("Config").InsertOne(r.context, config)
		if err != nil {
			return config, err
		}
		return config, nil
	}
	err = client.Database(r.dbName).Collection("Config").FindOne(r.context, bson.M{}).Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
func (r *MongoRepository) SaveConfig(config domain.Config) (domain.Config, error) {

	client := r.client
	_, err := client.Database(r.dbName).Collection("Config").UpdateOne(
		r.context,
		bson.M{},
		bson.D{
			{"$set", bson.D{{"heighttx", config.HeightTx}, {"heightvalid", config.HeightValid}}},
		},
	)
	if err != nil {
		return config, err
	}
	return config, nil

}
func (r *MongoRepository) SaveConfirmedTx(tx *domain.TransactionDetails, collectionName string) (bool, error) {
	client := r.client

	result, err := client.Database(r.dbName).Collection(collectionName).UpdateOne(
		r.context,
		bson.M{"confirmtxhash": tx.Hash},
		bson.D{
			{"$set", bson.D{{"confimationonledger", true}}},
		},
	)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount != 0, nil
}
