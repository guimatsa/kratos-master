ALTER TABLE "selfservice_login_request_methods" RENAME TO "selfservice_login_flow_methods";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "selfservice_login_requests" RENAME TO "selfservice_login_flows";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "selfservice_registration_request_methods" RENAME TO "selfservice_registration_flow_methods";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "selfservice_registration_requests" RENAME TO "selfservice_registration_flows";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "selfservice_settings_request_methods" RENAME TO "selfservice_settings_flow_methods";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "selfservice_settings_requests" RENAME TO "selfservice_settings_flows";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "selfservice_recovery_request_methods" RENAME TO "selfservice_recovery_flow_methods";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "selfservice_recovery_requests" RENAME TO "selfservice_recovery_flows";COMMIT TRANSACTION;BEGIN TRANSACTION;
ALTER TABLE "selfservice_verification_requests" RENAME TO "selfservice_verification_flows";COMMIT TRANSACTION;BEGIN TRANSACTION;