# Tech Design
Tech design document is to solidate as much implementation details as possible. It should cover macro system architecture chart, data structure schema design, database table design, logic flow chart and other UMLs. The most important thing is to help developers sort out their design and implementation ideas before actually writing code. Tech design documents can also serve as a reference when reviewing a project in the future.
# Overall Architecute
Considering the system usage scale, complex desgin won't be introduced into this project. 

![Architecute](image/architecture.png)

The Management Service is main component of this system. Considering the project implementation and debugging difficulty, SQLite as the database is enough. As there is a requirement in the requirement document, to update the course module status automatically, a cron job Course Status Job component is needed.

<mark>**Highlight**: In actual projects, course selection is a very complex process and it is a typical high-concurrency scenario in a short period of time. To solve this, not only the optimization of code logic level, but also the introduction of new middleware components, deployment of service instances and many other aspects will be taken into consideration. However, these contents are not the focus of this project, so the architectural design has been extremely simplified.</mark>

# Database Table Design
User Table:
```
CREATE TABLE IF NOT EXISTS user_tab (
    user_id INTEGER NOT NULL PRIMARY KEY,
    password VARCHAR(256) NOT NULL
);
```
College Table:
```
CREATE TABLE IF NOT EXISTS college_tab (
    college_id VARCHAR(256) NOT NULL PRIMARY KEY,
    college_name VARCHAR(256) NOT NULL
);
```
Role Table:
student, professor and administartor have many overlapped fields so that one unified table can be used to store all of them.

status: Normal, Graduated, Suspension, Retired
```
CREATE TABLE IF NOT EXISTS role_tab (
    role_id VARCHAR(256) NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    college_id VARCHAR(256) NOT NULL,
    name VARCHAR(256) NOT NULL,
    gender INTEGER,
    type INTEGER,
    email VARCHAR(256),
    grade INTEGER,
    enrollment_year INTEGER,
    status INTEGER,
    FOREIGEN KEY (user_id) REFERENCES user_tab (user_id),
    FOREIGEN KEY (college_id) REFERENCES college_tab (college_id)
);
```
Semester Table
```
CREATE TABLE IF NOT EXISTS semester_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    type INTEGER,
    semester VARCHAR(256) UNIQUE,
    start_time INTEGER,
    end_time INTEGER
);
```
Course Table
```
CREATE TABLE IF NOT EXISTS course_tab (
    course_id VARCHAR(256) NOT NULL PRIMARY KEY,
    course_name VARCHAR(256) NOT NULL,
    college_id VARCHAR(256) NOT NULL,
    credit INTEGER,
    brief TEXT,
    FOREIGEN KEY(college_id) REFERENCES college_tab(college_id)
);
```
Course Module Table

The format of score_ratio is "[{"type":1, "ratio":0.1},{"type":1, "ratio":0.1}, {"type":1, "ratio":0.1}, {"type":2, "ratio":0.2}, {"type":3, "ratio":0.2}, {"type":4, "ratio":0.3}]", 1: assignment, 2: quiz or midterm exam, 3: project, 4: final exam.

status: 1: Selection In Progress, 2: Normal Teaching, 3: Course Ended, 4: Canceled, 5: Reviewing.
```
CREATE TABLE IF NOT EXISTS course_module_tab (
    course_module_id TEXT NOT NULL PRIMARY KEY,
    course_id VARCHAR(256) NOT NULL,
    professor_id VARCHAR(256) NOT NULL,
    ta_id VARCHAR(256),
    semester VARCHAR(256) NOT NULL,
    classroom VARCHAR(256),
    class_period_start VARCHAR(256),
    class_period_end VARCHAR(256),
    duration INTEGER,
    course_capacity INTEGER,
    min_stu_num INTEGER,
    score_ratio TEXT,
    status INTEGER,
    FOREIGEN KEY(course_id) REFERENCES course_tab(course_id),
    FOREIGEN KEY(professor_id) REFERENCES role_tab(role_id),
    FOREIGEN KEY(ta_id) REFERENCES role_tab(role_id)
);
```
Course Module Student Table, to store the course selection data.
status is to mark the selection status, Selecting, Enrolled, Selection Failed, Ended, Failed.

Scores: "[100, 100, 100, 80, 90, 85]"
```
CREATE TABLE IF NOT EXISTS course_module_stu_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    course_module_id TEXT NOT NULL PRIMARY KEY,
    course_id VARCHAR(256) NOT NULL,
    stu_id VARCHAR(256) NOT NULL,
    scores TEXT,
    final_score INTEGER,
    status INTEGER,
    FOREIGEN KEY(course_module_id) REFERENCES course_module_tab(course_module_id),
    FOREIGEN KEY(course_id) REFERENCES course_tab(course_id),
    FOREIGEN KEY(stu_id) REFERENCES role_tab(role_id)
);
```
Assignment/Project/Quiz/Exam Table, the published assignment table.
Type: 1: assignment, 2:quiz, 3: project, 4: final exam. 
```
CREATE TABLE IF NOT EXISTS assignment_exam_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    course_module_id TEXT NOT NULL PRIMARY KEY,
    professor_id VARCHAR(256) NOT NULL,
    type INTEGER NOT NULL, 
    title VARCHAR(256) NOT NULL,
    content TEXT,
    publish_time INTEGER,
    deadline INTEGER,
    exam_start VARCHAR(256),
    exam_end VARCHAR(256),
    exam_room VARCHAR(256),
    FOREIGEN KEY(course_module_id) REFERENCES course_module_tab(course_module_id),
    FOREIGEN KEY(professor_id) REFERENCES role_tab(role_id)
);
```
Student Assignment/Project/Quiz/Exam Table, the table to store the students' submission of assignmeny/project, this table also stores the students' scores.
score_status: 1: Waiting, 2: Submitted, 3: Confirmed.
```
CREATE TABLE IF NOT EXISTS stu_assignment_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    assignment_id INTEGER,
    stu_id VARCHAR(256) NOT NULL,
    type INTEGER,
    last_submit_time INTEGER,
    content TEXT,
    score INTEGER,
    score_status INTEGER,
    FOREIGEN KEY(assignment_id) REFERENCES assignment_exam_tab(id)
);
```
Notification Sender Table
receiver_features: a string list, each element represent a feature.
is_all_users, is_all_students, is_all_normal_students, is_all_professors, is_all_normal_professors, course_id, course_module_id, receiver_role_ids

receiver_role_ids is also a list, like ["ST20150100126", "FA19980320082"]
```
CREATE TABLE IF NOT EXISTS noti_sender_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    sender_id VARCHAR(256) NOT NULL,
    title VARCHAR(256),
    content TEXT,
    send_time INTEGER,
    receiver_features TEXT,
    FOREIGEN KEY(sender_id) REFERENCES role_tab(role_id)
);
```
Notification Receiver Table status: 1: Unread, 2: Have Read, 3: Deleted
```
CREATE TABLE IF NOT EXISTS noti_sender_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    sender_id VARCHAR(256) NOT NULL,
    title VARCHAR(256),
    content TEXT,
    send_time INTEGER,
    receiver_id VARCHAR(256) NOT NULL,
    status INTEGER,
    FOREIGEN KEY(sender_id) REFERENCES role_tab(role_id),
    FOREIGEN KEY(receiver_id) REFERENCES role_tab(role_id)
);
```