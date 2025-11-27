-- Заполнение таблицы user
INSERT INTO users (login, hashed_password, is_moderator) VALUES
('admin', 'hashed_admin_password', true),
('user1', 'hashed_user1_password', false);

-- Заполнение таблицы Component
INSERT INTO components (title, type, mtbf, mttr, available, img, description, is_deleted) VALUES
('Одиночный сервер базы данных', 'Одиночный сервер', 8760, 2, 0.999771, 'http://localhost:9000/failivercheck/svg/database.svg', 'Один экземпляр сервера, содержащий базу данных.  Подвержен простоям при отказах и обслуживании.', false),
('Балансировщик нагрузки (активный/пассивный)', 'Балансировщик', 43800, 4, 0.999909, 'http://localhost:9000/failivercheck/svg/loader.svg', 'Два балансировщика нагрузки, один активный, другой пассивный. При отказе активного, пассивный автоматически занимает его место.', false),
('Балансировщик нагрузки (геораспределенный)', 'Балансировщик', 52560, 3, 0.999943, 'http://localhost:9000/failivercheck/svg/loader.svg', 'Три балансировщика нагрузки, размещенных в разных географических регионах. Распределяют нагрузку и обеспечивают отказоустойчивость даже при выходе из строя целого региона.', false),
('Балансировщик нагрузки (VRRP)', 'Балансировщик', 43800, 4, 0.999909, 'http://localhost:9000/failivercheck/svg/loader.svg', 'Два балансировщика нагрузки, использующие VRRP (Virtual Router Redundancy Protocol) для обеспечения отказоустойчивости. Один активен, другой - в режиме ожидания.', false),
('Веб-сервер (кластер)', 'Сервер', 8760, 2, 0.999771, 'http://localhost:9000/failivercheck/svg/cluster.svg', 'Несколько веб-серверов, работающих вместе для обработки запросов. Если один сервер выходит из строя, другие продолжают работать.', false),
('Кэширующий сервер (распределенный)', 'Сервер', 26280, 3, 0.999885, 'http://localhost:9000/failivercheck/svg/database.svg', 'Три кэширующих сервера, объединенных в кластер. Данные распределены между серверами, обеспечивая отказоустойчивость и масштабируемость.', false),
('Очередь сообщений (кластер)', 'Сервис', 35040, 4, 0.999886, 'http://localhost:9000/failivercheck/svg/database.svg', 'Четыре сервера очереди сообщений, работающих в кластере. Обеспечивают высокую доступность и надежность доставки сообщений.', false);

-- Заполнение таблицы SystemCalculation
INSERT INTO system_calculations (system_name, available_calculation, user_id, status, date_created)
VALUES ('System 1', 98.1, 1, 'DRAFT', NOW());

-- Заполнение связующей таблицы ComponentsToSystemCalc
INSERT INTO components_to_system_calcs (component_id, system_calculation_id, replication_count) VALUES
(1, 1, 1),
(2, 1, 2),
(3, 1, 4);