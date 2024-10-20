-- admin-password
-- user-password
INSERT INTO user (email, password)
VALUES ('admin@example.com', '$2a$14$wLX1x9J5dAwazUvlFT8Or.fsZwUCYGokUYgkPPabDNXF6LXFxXtYS'), -- pwd = admin
       ('user@example.com', '$2a$14$K5GN7g.pA8Inumq7Dybpm.nvawQ0NghLlNUy1t7QxW4Rf./GtOXR6'); -- pwd = user

INSERT INTO event (title, description, location, date_time, user_id)
VALUES ('Event 1', 'Description for Event 1', 'Location 1', '2024-10-18 10:00:00', 1),
       ('Event 2', 'Description for Event 2', 'Location 2', '2024-10-19 11:00:00', 1),
       ('Event 3', 'Description for Event 3', 'Location 3', '2024-10-20 12:00:00', 1),
       ('Event 4', 'Description for Event 4', 'Location 4', '2024-10-21 13:00:00', 1),
       ('Event 5', 'Description for Event 5', 'Location 5', '2024-10-22 14:00:00', 1),
       ('Event 6', 'Description for Event 6', 'Location 6', '2024-10-23 15:00:00', 1),
       ('Event 7', 'Description for Event 7', 'Location 7', '2024-10-24 16:00:00', 1),
       ('Event 8', 'Description for Event 8', 'Location 8', '2024-10-25 17:00:00', 1),
       ('Event 9', 'Description for Event 9', 'Location 9', '2024-10-26 18:00:00', 1),
       ('Event 10', 'Description for Event 10', 'Location 10', '2024-10-27 19:00:00', 2);