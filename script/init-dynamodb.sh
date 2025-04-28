#!/bin/sh

TABLE_NAME="i.maker_Users"
REGION="ap-northeast-1"
ENDPOINT="http://dynamodb-circle:8000"

# テーブルが存在するかチェック
EXISTING_TABLE=$(aws dynamodb list-tables --region $REGION --endpoint-url $ENDPOINT --query "TableNames" --output json | grep -w "$TABLE_NAME")

if [ -z "$EXISTING_TABLE" ]; then
    echo "Creating table: $TABLE_NAME ..."
    aws dynamodb create-table --table-name $TABLE_NAME \
        --region $REGION \
        --endpoint-url $ENDPOINT \
        --attribute-definitions \
            AttributeName=user_id,AttributeType=S \
            AttributeName=created_at,AttributeType=S \
        --key-schema \
            AttributeName=user_id,KeyType=HASH \
            AttributeName=created_at,KeyType=RANGE \
        --provisioned-throughput ReadCapacityUnits=10,WriteCapacityUnits=10 \
        --no-cli-pager
    echo "Table $TABLE_NAME created successfully!"
else
    echo "Table $TABLE_NAME already exists. Skipping creation."
fi
