Table "distributors" {
  "distributor_id" INT [pk]
  "user_id" INT
}

Table "packages" {
  "package_id" INT [pk]
  "distributor_id" INT
  "origin" VARCHAR(255)
  "destination" VARCHAR(255)
  "type" VARCHAR(100)
  "est_value" float
  "offer_price" float
  "publication_date" DATE
  "modification_date" DATE
  "deadline_date" DATE
  "dimensiones" VARCHAR(50)
  "eu_palet" BOOLEAN
  "accepted" BOOLEAN
  "completed" BOOLEAN
}

Table "transporters" {
  "transporter_id" INT [pk]
  "user_id" INT
}

Table "vehicles" {
  "vehicle_id" INT [pk]
  "transporter_id" INT
  "plate" VARCHAR(15)
  "brand" VARCHAR(50)
  "model" VARCHAR(50)
  "capacity_kg" DECIMAL(10,2)
}

Table "orders" {
  "order_id" bigint [pk]
  "package_id" INT
  "transporter_id" INT
  "distribuitor_id" INT
  "order_price" float
  "modification_date" TIMESTAMP
  "est_arrival_date" TIMESTAMP
  "order_date" TIMESTAMP
  "order_status" VARCHAR(50)
}

Table "tracking" {
  "tracking_id" INT [pk]
  "order_id" INT
  "latitude" DECIMAL(10,6)
  "longitude" DECIMAL(10,6)
  "modification_date" TIMESTAMP
}

Table "documents" {
  "document_id" INT [pk]
  "user_id" INT
  "order_id" bigint
  "type" VARCHAR(255)
  "description" VARCHAR(255)
  "name" VARCHAR(255)
}

Table "users" {
  "user_id" INT [pk]
  "name" VARCHAR(255)
  "email" VARCHAR(255)
  "phone" VARCHAR(20)
  "address" VARCHAR(255)
  "type" VARCHAR(5)
  "vat_number" VARCHAR(30)
  "is_verified" BOOLEAN
  "modification_date" DATE
  "password" VARCHAR(255)
}

Table "audit_log" {
  "log_id" INT [pk, increment]
  "table_name" VARCHAR(255)
  "action" VARCHAR(10)
  "change_date" TIMESTAMP
  "user_id" INT
  "executed_sql" TEXT
}


Ref:"distributors"."distributor_id" < "packages"."distributor_id"

Ref:"transporters"."transporter_id" < "vehicles"."transporter_id"

Ref:"users"."user_id" < "documents"."user_id"


Ref: "transporters"."transporter_id" < "orders"."transporter_id"

Ref: "distributors"."distributor_id" < "orders"."distribuitor_id"

Ref: "users"."user_id" < "distributors"."user_id"

Ref: "users"."user_id" < "transporters"."user_id"

Ref: "tracking"."order_id" < "orders"."order_id"

Ref: "documents"."order_id" < "orders"."order_id"

Ref: "orders"."package_id" < "packages"."package_id"
