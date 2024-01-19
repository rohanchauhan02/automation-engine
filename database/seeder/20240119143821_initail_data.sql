-- +goose Up
-- +goose StatementBegin

INSERT INTO public.tasks
(title, description, priority, due_date, user_id, created_at, updated_at)
VALUES('', '', '', '', 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.users
("name", email, phone_number, address, created_at, updated_at)
VALUES('Jhon', 'jhon@gmail.com', '92112233', 'USA', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.users
("name", email, phone_number, address, created_at, updated_at)
VALUES('Jean', 'jean@gmail.com', '92121213', 'Canada', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.users
("name", email, phone_number, address, created_at, updated_at)
VALUES('Rohan', 'rohan@gmail.com', '925453245', 'New Delhi, India', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.users
("name", email, phone_number, address, created_at, updated_at)
VALUES('Aldi', 'aldi@gmail.com', '92567853', 'SA', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.users
("name", email, phone_number, address, created_at, updated_at)
VALUES('Devid', 'devid@gmail.com', '928797657', 'Australia', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.users
("name", email, phone_number, address, created_at, updated_at)
VALUES('Mike', 'mike@gmail.com', '92099787', 'China', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
