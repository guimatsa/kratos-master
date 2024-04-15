ALTER TABLE "identity_recovery_tokens" ADD COLUMN "expires_at" timestamp NOT NULL DEFAULT '2000-01-01 00:00:00';COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "identity_recovery_tokens" ADD COLUMN "issued_at" timestamp NOT NULL DEFAULT '2000-01-01 00:00:00';COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "identity_recovery_tokens" DROP CONSTRAINT "identity_recovery_tokens_selfservice_recovery_requests_id_fk";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "identity_recovery_tokens" RENAME COLUMN "selfservice_recovery_flow_id" TO "_selfservice_recovery_flow_id_tmp";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "identity_recovery_tokens" ADD COLUMN "selfservice_recovery_flow_id" UUID;COMMIT TRANSACTION;BEGIN TRANSACTION;
UPDATE "identity_recovery_tokens" SET "selfservice_recovery_flow_id" = "_selfservice_recovery_flow_id_tmp";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "identity_recovery_tokens" DROP COLUMN "_selfservice_recovery_flow_id_tmp";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "identity_recovery_tokens" ADD CONSTRAINT "identity_recovery_tokens_selfservice_recovery_requests_id_fk" FOREIGN KEY ("selfservice_recovery_flow_id") REFERENCES "selfservice_recovery_flows" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;COMMIT TRANSACTION;BEGIN TRANSACTION;