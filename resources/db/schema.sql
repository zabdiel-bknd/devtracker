-- schema.sql

-- Projects table: Stores the high-level projects
CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tasks table: Stores individual tasks associated with a project
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    
    -- Priority constraint: Simulating an Enum
    priority VARCHAR(10) NOT NULL CHECK (priority IN ('LOW', 'MEDIUM', 'HIGH')),
    
    -- Status constraint: Simulating an Enum
    status VARCHAR(10) NOT NULL CHECK (status IN ('TODO', 'DOING', 'DONE')),
    
    project_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign Key: If a project is deleted, delete its tasks (Cascade)
    CONSTRAINT fk_project
        FOREIGN KEY(project_id) 
        REFERENCES projects(id)
        ON DELETE CASCADE 
);

-- Seed data: Initial data to verify connection later
INSERT INTO projects (name, description) VALUES ('Learn Go', 'Backend engineering workshop in Go');