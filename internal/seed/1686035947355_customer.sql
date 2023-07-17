-- INSERT INTO "customers" (name, ktp_number, ktp_image,is_corporate,corporate_name,account_status,request_status) VALUES ('Bayu Sarifudin', '3374123344441234', 'ktp-img-bayu', true, 'bayucorp', 'active', 'packaging');

-- INSERT INTO customers (email, password, phone_number, role_id, type, account_status, corporate_name, is_corporate, ktp_image, ktp_number, name, nib, npwp_image, npwp_number, request_status) VALUES ('yohanes@gmail.com', 'amansaja', '087234555444', 1, '', 'non-active', 'Tidak ada', false, 'ktp-img-1', '3374123344447766', 'Yohanes Adi Prayoga', '111000333999', 'npwp-img-1', '555321', 'packaging');

UPDATE customers SET is_corporate = true WHERE id = 5;