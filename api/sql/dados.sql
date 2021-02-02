insert into usuarios(nome, nick, email, senha)
values
("usuario 1", "usuario1", "usuario1@gmail.com", "$2a$10$8a4GeVROcznTOPQ2sbYVnOBDeDFMDCRHdyc0LgwyZac3OjaLXZkQC"),
("usuario 2", "usuario2", "usuario2@gmail.com", "$2a$10$8a4GeVROcznTOPQ2sbYVnOBDeDFMDCRHdyc0LgwyZac3OjaLXZkQC"),
("usuario 3", "usuario3", "usuario3@gmail.com", "$2a$10$8a4GeVROcznTOPQ2sbYVnOBDeDFMDCRHdyc0LgwyZac3OjaLXZkQC");

insert into seguidores(usuario_id, seguidor_id)
values
(1,2),
(3,1),
(1,3);

INSERT INTO publicacoes (titulo, conteudo, autor_id)
values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3!", 3);