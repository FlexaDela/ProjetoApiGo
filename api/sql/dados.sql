insert into usuarios(nome, nick, email, senha)
values
("usuario 1", "usuario_1", "usuario1@gmail.com"),
("usuario 2", "usuario_2", "usuario2@gmail.com"),
("usuario 3", "usuario_3", "usuario3@gmail.com");

insert into seguidores(usuario_id, seguidor_id)
values
(1,2),
(3,1),
(1,3);

insert into publicacoes(titulo, conteudo, autor_id)
values("Publicação do usuario 1",1),
values("Publicação do usuario 1",2),
values("Publicação do usuario 1",3),