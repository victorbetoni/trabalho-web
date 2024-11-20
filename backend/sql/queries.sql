-- name: CreateProfessor :exec
INSERT INTO Professores (CPF, Nome, Formacao, Telefone, Rua, Bairro, CEP, Cidade, Numero) VALUES (?,?,?,?,?,?,?,?,?);

-- name: CreateAula :exec
INSERT INTO Aulas (ID, MateriaID, Professor, CargaHorariaMinutos, Data) VALUES (?,?,?,?,?);

-- name: CreateMateria :exec
INSERT INTO Materias (ID, Nome, CargaHorariaMinutos) VALUES (?,?,?);

-- name: CreateAluno :exec
INSERT INTO Alunos (GRR, Nome) VALUES (?,?);

-- name: CreatePresenca :exec
INSERT INTO Presencas (AulaID, AlunoGRR) VALUES (?,?);

-- name: FindProfessores :many
SELECT * FROM Professores
WHERE (sqlc.arg(cpf) = "" OR CPF LIKE sqlc.arg(cpf)) 
AND (sqlc.arg(nome) = "" OR Nome LIKE sqlc.arg(nome))
LIMIT ? OFFSET ?;

-- name: FindAula :many
SELECT * FROM Aulas
INNER JOIN Materias ON Materias.ID = MateriaID
WHERE (sqlc.arg(materianome) = "" OR Materias.Nome LIKE sqlc.arg(materianome)) 
AND (sqlc.arg(materiaid) = "" OR Materias.ID LIKE sqlc.arg(materiaid)) 
AND (sqlc.arg(aulaid) = "" OR Aulas.ID LIKE sqlc.arg(aulaid))
LIMIT ? OFFSET ?;

-- name: FindAluno :many
SELECT * FROM Alunos
WHERE (sqlc.arg(nome) = "" OR Nome LIKE sqlc.arg(nome)) 
AND (sqlc.arg(grr) = "" OR GRR LIKE sqlc.arg(grr)) 
LIMIT ? OFFSET ?;

-- name: FindMateria :many
SELECT * FROM Materias
WHERE (sqlc.arg(nome) = "" OR Nome LIKE sqlc.arg(nome)) 
AND (sqlc.arg(id) = "" OR ID LIKE sqlc.arg(id)) 
LIMIT ? OFFSET ?;

-- name: UpdateProfessor :exec
UPDATE Professores SET Nome = ? WHERE CPF = ?;

-- name: UpdateAluno :exec
UPDATE Alunos SET Nome = ? WHERE GRR = ?;

-- name: UpdateMateria :exec
UPDATE Materias SET Nome = ?, CargaHorariaMinutos = ? WHERE ID = ?;

-- name: UpdateAula :exec
UPDATE Aulas SET MateriaID = ?, Professor = ?, Data = ? WHERE ID = ?;

-- name: DeleteProfessor :exec
DELETE FROM Professores WHERE CPF = ?;

-- name: DeleteAluno :exec
DELETE FROM Alunos WHERE GRR = ?;

-- name: DeleteMateria :exec
DELETE FROM Materias WHERE ID = ?;

-- name: DeleteAula :exec
DELETE FROM Aulas WHERE ID = ?;

-- name: FindLogin :one
SELECT Senha FROM Login WHERE Username = ?;