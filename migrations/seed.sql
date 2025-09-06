DELETE FROM StudentPrograms;
DELETE FROM Enrollments;
DELETE FROM Students;
DELETE FROM Programs;
DELETE FROM Program_Types;
DELETE FROM Departments;
DELETE FROM Units;
DELETE FROM Unit_Types;
DELETE FROM Courses;
DELETE FROM Statuses;

PRAGMA foreign_keys = ON;

INSERT INTO Unit_Types (unit_type_id, unit_type_name) VALUES
(1, 'Academic Unit');

INSERT INTO Units (unit_id, unit_type_id, unit_name, unit_acronym) VALUES
(1, 1, 'College of ISC', 'CISC');

INSERT INTO Departments (dept_id, unit_id, dept_name) VALUES
(10, 1, 'IT Department'),
(11, 1, 'CS Department'),
(12, 1, 'IS Department');

INSERT INTO Program_Types (prog_type_id, prog_type_name) VALUES
(1, 'Undergraduate');

INSERT INTO Programs (program_id, dept_id, prog_type_id, program_name, program_reg_years) VALUES
(200, 10, 1, 'BS Information Technology', 4),
(201, 11, 1, 'BS Computer Science', 4),
(202, 12, 1, 'BS Information Systems', 4);

INSERT INTO Students (st_id, fname, mname, lname, sex, birthdate, year) VALUES
(1001, 'Naruto', 'U.', 'Uzumaki', 'M', '2004-01-15', 2),
(1002, 'Sakura', 'H.', 'Haruno', 'F', '2003-05-22', 3),
(1003, 'Sasuke', 'U.', 'Uchiha', 'M', '2002-08-30', 4),
(1004, 'Hinata', 'H.', 'Hyuga', 'F', '2004-02-10', 2),
(1005, 'Luffy', 'D.', 'Monkey', 'M', '2003-03-25', 3),
(1006, 'Nami', 'S.', 'Navigator', 'F', '2004-06-18', 2),
(1007, 'Zoro', 'R.', 'Roronoa', 'M', '2002-09-12', 4),
(1008, 'Goku', 'S.', 'Son', 'M', '2003-11-02', 3),
(1009, 'Vegeta', 'P.', 'Prince', 'M', '2004-04-20', 2);

INSERT INTO StudentPrograms (st_id, program_id) VALUES
(1001, 200),
(1002, 200),
(1003, 201),
(1004, 202),
(1005, 200),
(1006, 202),
(1007, 201),
(1008, 200),
(1009, 202);

INSERT INTO Courses (course_id, course_name, course_units) VALUES
(301, 'Database Systems', 3),
(302, 'Software Engineering', 3),
(303, 'Web Development', 3),
(304, 'Data Structures', 4),
(305, 'Business Analytics', 3),
(306, 'Computer Networks', 3),
(307, 'Systems Analysis', 3),
(308, 'Artificial Intelligence', 4),
(309, 'Operating Systems', 3),
(310, 'IT Project Management', 3);

INSERT INTO Statuses (status_id, status_name) VALUES
(1, 'Enrolled'),
(2, 'Completed'),
(3, 'Dropped');

INSERT INTO Enrollments (enrollment_id, st_id, course_id, status_id, academic_year, academic_semester, midterms_grade, finals_grade) VALUES
(1, 1001, 301, 1, 2024, '1', 88, 90),
(2, 1001, 302, 2, 2024, '1', 85, 89),
(3, 1002, 303, 1, 2024, '1', 87, 91),
(4, 1003, 304, 3, 2024, '1', 70, 0),
(5, 1004, 305, 1, 2024, '1', 92, 95),
(6, 1005, 306, 1, 2024, '1', 86, 90),
(7, 1006, 307, 1, 2024, '1', 89, 93),
(8, 1007, 308, 2, 2024, '1', 91, 94),
(9, 1008, 309, 1, 2024, '1', 84, 88),
(10, 1009, 310, 1, 2024, '1', 90, 92);