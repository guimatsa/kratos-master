CREATE UNIQUE INDEX "identity_verifiable_addresses_status_via_uq_idx" ON "identity_verifiable_addresses" (nid, via, value);