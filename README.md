![image](https://github.com/mnguyen081002/krakend-keycloak-authorize/assets/76799726/ab83b6a8-0044-44ab-bb8f-7e70eb94869f)# Krakend Keycloak Authorize

## Các bước để bắt đầu

### 1. Clone this repo

- `git clone https://github.com/mnguyen081002/krakend-keycloak-authorize`

### 2. Run helm (phải bật ingress addons để test postman hoặc có thể thủ công forward port)

- `cd krakend-keycloak-authorize;eval $(minikube docker-env);docker build ./order-service -t order-service; docker build ./customer-service -t customer-service; docker build ./krakend -t krakend; helm install krakend-keycloak-authorize helm`

### 3. Cấu hình keycloak

### Tạo client

1. Mở url http://keycloak.test/ (nếu xài ingress). Đăng nhập account admin/admin
2. Vào menu **Clients**, tạo client id `test` sau đó ấn **Next**.
   ![image](https://github.com/mnguyen081002/krakend-keycloak-authorize/assets/76799726/8a862173-5c62-4d0d-bb6b-47aa45e63915)
3. Bật **Client authentication** và **Authorization** sau đó ấn **Next**.
   ![image](https://github.com/mnguyen081002/krakend-keycloak-authorize/assets/76799726/f0063e5d-5de7-4234-bbc9-0844894a9024)
4. Tiếp tục ấn **Save**.![image](https://github.com/mnguyen081002/krakend-keycloak-authorize/assets/76799726/1ed69979-8ace-434f-ba11-91695889a2bf)

### Tạo client scope

1. Chọn vào menu **Client scopes**, chọn **Create client scope**
2. Nhập name **order:view** và chọn **Save**![image](https://github.com/mnguyen081002/krakend-keycloak-authorize/assets/76799726/e41d5bb5-6726-4bad-811c-a7b337eb40e9)

### Thêm client scope và client

1. Chọn menu **Clients**, ấn vào client **test**![image](https://github.com/mnguyen081002/krakend-keycloak-authorize/assets/76799726/40260ba1-fa5e-48cd-ba0b-f7ed13e6c673)
2. Ấn vào **Client scopes** và chọn **Add client scope**![image](https://github.com/mnguyen081002/krakend-keycloak-authorize/assets/76799726/0274cc44-2561-4fea-97a6-a84f8e04df58)
3. Chọn **order:view** và ở phần **Add** chọn **Optional**![image](https://github.com/mnguyen081002/krakend-keycloak-authorize/assets/76799726/5852069f-9385-491a-9615-1a1781592d6f)

### Tạo role customer

1. Vào menu **Realm roles**, ấn **Create role** và nhập name **customer** sau đó **Save**

### Tạo group

1. Vào menu **Groups** và ấn **Create group**, nhập name **user** sau đó **Create**
2. Ấn vào group **user** vừa tạo, sau đó chọn **Role mapping** và assign role **customer**

### Tạo user

1. Vào menu **Users** và ấn **Add user** với username là `2051120273` (tùy chọn)
2. Sau đó chọn vào user vừa tạo, chọn **Role mapping** và chọn tiếp **Assign role**, sau đó assign role **customer** và ấn **Assign**
3. Chọn vào **Credentials** và set password `123456` (tùy chọn)
4. Chọn lại vào user vừa tạo và tắt **require action** ![image](https://github.com/mnguyen081002/krakend-keycloak-authorize/assets/76799726/4a39b9a8-49bb-48f0-bddc-cde6be4587c2)

### 4. Test với postman

### Chuẩn bị:

1. Import **KeyCloak.postman_collection.json** trong repo vào Postman
2. Vào **Client** -> Chọn client `test` -> Chọn **Credential** -> Copy **client secret** để và thay đổi **client_secret** trong body ở api **Login**

### Test:

1. Trường hợp 403

- **Tắt scope ở body** và ấn **Send**
- Copy `access_token` và thử với **api orders** và **profile**
- Result: Order 403 và Profile 200

2. Trường hợp 200

- **Bật scope ở body** và ấn **Send**
- Copy `access_token` và thử với **api orders** và **profile**
- Result: Order 200 và Profile 200
