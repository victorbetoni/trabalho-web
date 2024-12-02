-- name: CreateProfessor :exec
INSERT INTO Professores (CPF, Nome, Formacao, Telefone, Rua, Bairro, CEP, Cidade, Numero) VALUES (?,?,?,?,?,?,?,?,?);

-- name: CreateAula :exec
INSERT INTO Aulas (ID, Data, Professor, CargaHorariaMinutos, Titulo) VALUES (?,?,?,?,?);

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
INNER JOIN Professores ON Professores.CPF = Aulas.Professor
WHERE (sqlc.arg(professor) = "" OR Professores.CPF LIKE sqlc.arg(professor)) 
ORDER BY Aulas.Data DESC LIMIT ? OFFSET ?;

-- name: FindAluno :many
SELECT * FROM Alunos
WHERE (sqlc.arg(nome) = "" OR Nome LIKE sqlc.arg(nome)) 
AND (sqlc.arg(grr) = "" OR GRR LIKE sqlc.arg(grr)) 
LIMIT ? OFFSET ?;

-- name: UpdateProfessor :exec
UPDATE Professores SET Nome = ?, Formacao = ?, Telefone = ?, CEP = ?, Rua = ?, Bairro = ?, Cidade = ?, Numero = ? WHERE CPF = ?;

-- name: UpdateAluno :exec
UPDATE Alunos SET Nome = ? WHERE GRR = ?;

-- name: DeleteProfessor :exec
DELETE FROM Professores WHERE CPF = ?;

-- name: DeleteAluno :exec
DELETE FROM Alunos WHERE GRR = ?;

-- name: DeleteAula :exec
DELETE FROM Aulas WHERE ID = ?;

-- name: FindLogin :one
SELECT Senha FROM Login WHERE Username = ?;

-- name: AulasDadas :one
SELECT COUNT(*) FROM Aulas WHERE Professor = ?;