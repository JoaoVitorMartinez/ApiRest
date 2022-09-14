create table jogos(
    id serial primary key,
    nome varchar,
    descricao varchar,
    preco decimal,
    nota decimal
);

INSERT INTO jogos(nome, descricao, preco, nota) VALUES
('Far Cry 6', 'Livre', 200, 10),
('Far cry 5', 'Livre', 100, 9);