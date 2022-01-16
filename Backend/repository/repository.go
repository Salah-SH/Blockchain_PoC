package repository

import (
	"context"
	"log"
	"os"

	"github.com/cagnotteApp/Backend/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	context context.Context
	client  *mongo.Client
	dbName  string
}
type MongoDB struct {
	client *mongo.Client
	dbName string
}

var client *mongo.Client

type DB interface{}

func NewDB() (DB, error) {

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "mongodb://localhost:27017/Blockchain"

	}
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return MongoDB{
		client: client,
		dbName: "Blockchain",
	}, nil
}

func New(ctx context.Context, db DB) Repository {

	mongoDB := db.(MongoDB)
	return &MongoRepository{
		context: ctx,
		client:  mongoDB.client,
		dbName:  mongoDB.dbName,
	}
}

type Repository interface {
	RegisterPaiement(accAddr string, paiementoken string, cagnotteName string, amount string) error
	UpdatePaiement(accAddr string, paiementoken string, status bool) error
	GetAdminCreds() (domain.AdminCreds, error)
	RegisterAdminCreds(admin *domain.AdminCreds) (*domain.AdminCreds, error)
	SaveUser(user *domain.User) (*domain.User, error)
	FindUser(accAddr string) (*domain.User, error)
	FindCagnotte(name string) (*domain.Cagnotte, error)
	SaveCagnotte(cagnotte *domain.Cagnotte) (*domain.Cagnotte, error)
	UpdateCagnotte(cagnotte *domain.Cagnotte) (bool, error)
	SaveAddTx(txs *domain.TransactionDetails, collectionName string) (bool, error)
	FindTx(collectionName string, status ...string) ([]domain.TransactionDetails, error)
	UpdateDB(txs domain.TransactionDetails, success bool, collectionName string) (bool, error)
	UpdateValidationDB(txs domain.TransactionDetails, success bool, collectionName string, validationTxHash string) (bool, error)
	GetConfig() (domain.Config, error)
	SaveConfig(config domain.Config) (domain.Config, error)
	SaveConfirmedTx(txs *domain.TransactionDetails, collectionName string) (bool, error)
}
