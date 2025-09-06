CREATE TABLE IF NOT EXISTS Students (
    st_id INTEGER PRIMARY KEY AUTOINCREMENT,
    fname TEXT NOT NULL,
    mname TEXT,
    lname TEXT NOT NULL,
    sex TEXT CHECK (sex IN ('M','F')),
    birthdate TEXT,
    year INT
);

CREATE TABLE IF NOT EXISTS Unit_Types (
    unit_type_id INTEGER PRIMARY KEY AUTOINCREMENT,
    unit_type_name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Units (
    unit_id INTEGER PRIMARY KEY AUTOINCREMENT,
    unit_type_id INT NOT NULL,
    unit_name TEXT NOT NULL UNIQUE,
    unit_acronym TEXT UNIQUE,
    FOREIGN KEY (unit_type_id) REFERENCES Unit_Types(unit_type_id)
);

CREATE TABLE IF NOT EXISTS Departments (
    dept_id INTEGER PRIMARY KEY AUTOINCREMENT,
    unit_id INT NOT NULL,
    dept_name TEXT NOT NULL UNIQUE,
    FOREIGN KEY (unit_id) REFERENCES Units(unit_id)
);

CREATE TABLE IF NOT EXISTS Program_Types (
    prog_type_id INTEGER PRIMARY KEY AUTOINCREMENT,
    prog_type_name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Programs (
    program_id INTEGER PRIMARY KEY AUTOINCREMENT,
    dept_id INT NOT NULL,
    prog_type_id INT NOT NULL,
    program_name TEXT NOT NULL UNIQUE,
    program_reg_years INT,
    FOREIGN KEY (dept_id) REFERENCES Departments(dept_id),
    FOREIGN KEY (prog_type_id) REFERENCES Program_Types(prog_type_id)
);

CREATE TABLE IF NOT EXISTS StudentPrograms (
    st_id INT NOT NULL,
    program_id INT NOT NULL,
    FOREIGN KEY (st_id) REFERENCES Students(st_id),
    FOREIGN KEY (program_id) REFERENCES Programs(program_id)
);

CREATE TABLE IF NOT EXISTS Courses (
    course_id INTEGER PRIMARY KEY AUTOINCREMENT,
    course_name TEXT NOT NULL UNIQUE,
    course_units INT
);

CREATE TABLE IF NOT EXISTS Statuses (
    status_id INTEGER PRIMARY KEY AUTOINCREMENT,
    status_name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Enrollments (
    enrollment_id INTEGER PRIMARY KEY AUTOINCREMENT,
    st_id INT NOT NULL,
    course_id INT NOT NULL,
    status_id INT NOT NULL,
    academic_year INT NOT NULL,
    academic_semester TEXT CHECK (academic_semester IN ('1','2','M')) NOT NULL,
    midterms_grade INT CHECK (midterms_grade BETWEEN 0 AND 100),
    finals_grade INT CHECK (finals_grade BETWEEN 0 AND 100),
    FOREIGN KEY (st_id) REFERENCES Students(st_id),
    FOREIGN KEY (course_id) REFERENCES Courses(course_id),
    FOREIGN KEY (status_id) REFERENCES Statuses(status_id)
);