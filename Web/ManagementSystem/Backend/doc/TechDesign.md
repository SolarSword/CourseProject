# Tech Design
Tech design document is to solidate as much implementation details as possible. It should cover macro system architecture chart, data structure schema design, database table design, logic flow chart and other UMLs. The most important thing is to help developers sort out their design and implementation ideas before actually writing code. Tech design documents can also serve as a reference when reviewing a project in the future.
# Overall Architecute
Considering the system usage scale, complex desgin won't be introduced into this project. 

![Architecute](image/architecture.png)

The Management Service is main component of this system. Considering the project implementation and debugging difficulty, SQLite as the database is enough. As there is a requirement in the requirement document, to update the course module status automatically, a cron job Course Status Job component is needed.

<mark>**Highlight**: In actual projects, course selection is a very complex process and it is a typical high-concurrency scenario in a short period of time. To solve this, not only the optimization of code logic level, but also the introduction of new middleware components, deployment of service instances and many other aspects will be taken into consideration. However, these contents are not the focus of this project, so the architectural design has been extremely simplified.</mark>

By right, there should be Authentication conponent as the requirement document mentioned different permissions for different roles. But considering the project complexity, this part won't be introduced. I will put it in a separated project instead. <mark>So in this project, all the permission check will be very naive and straightforward.</mark>

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

Student = 1, Professor = 2, Administrator = 3. 

male = 1, female = 2. Freshman = 1, Sophomore = 2, Junior = 3, Senior = 4, 0 is used for professor or administrator. 

status: Normal = 1, Graduated = 2, Suspension = 3, Retired = 4
```
CREATE TABLE IF NOT EXISTS role_tab (
    role_id VARCHAR(256) NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES user_tab (user_id),
    college_id VARCHAR(256) NOT NULL REFERENCES college_tab (college_id),
    name VARCHAR(256) NOT NULL,
    gender INTEGER,
    type INTEGER,
    email VARCHAR(256),
    grade INTEGER,
    enrollment_year INTEGER,
    status INTEGER
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
    college_id VARCHAR(256) NOT NULL REFERENCES college_tab(college_id),
    credit INTEGER,
    brief TEXT
);
```
Course Module Table

The format of score_ratio is "[{"type":1, "ratio":0.1},{"type":1, "ratio":0.1}, {"type":1, "ratio":0.1}, {"type":2, "ratio":0.2}, {"type":3, "ratio":0.2}, {"type":4, "ratio":0.3}]", 1: assignment, 2: quiz or midterm exam, 3: project, 4: final exam.

status: 1: Selection In Progress, 2: Normal Teaching, 3: Course Ended, 4: Canceled, 5: Reviewing.
```
CREATE TABLE IF NOT EXISTS course_module_tab (
    course_module_id TEXT NOT NULL PRIMARY KEY,
    course_id VARCHAR(256) NOT NULL REFERENCES course_tab(course_id),
    professor_id VARCHAR(256) NOT NULL REFERENCES role_tab(role_id),
    ta_id VARCHAR(256) REFERENCES role_tab(role_id),
    semester VARCHAR(256) NOT NULL,
    classroom VARCHAR(256),
    class_period_start VARCHAR(256),
    class_period_end VARCHAR(256),
    duration INTEGER,
    course_capacity INTEGER,
    min_stu_num INTEGER,
    score_ratio TEXT,
    status INTEGER
);
```
Course Module Student Table, to store the course selection data.
status is to mark the selection status, Selecting, Enrolled, Selection Failed, Ended, Failed.

Scores: "[100, 100, 100, 80, 90, 85]"
```
CREATE TABLE IF NOT EXISTS course_module_stu_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    course_module_id TEXT NOT NULL REFERENCES course_module_tab(course_module_id),
    course_id VARCHAR(256) NOT NULL REFERENCES course_tab(course_id),
    stu_id VARCHAR(256) NOT NULL REFERENCES role_tab(role_id),
    scores TEXT,
    final_score INTEGER,
    status INTEGER
);
```
Assignment/Project/Quiz/Exam Table, the published assignment table.
Type: 1: assignment, 2:quiz, 3: project, 4: final exam. 
```
CREATE TABLE IF NOT EXISTS assignment_exam_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    course_module_id TEXT NOT NULL REFERENCES course_module_tab(course_module_id),
    professor_id VARCHAR(256) NOT NULL REFERENCES role_tab(role_id),
    type INTEGER NOT NULL, 
    title VARCHAR(256) NOT NULL,
    content TEXT,
    publish_time INTEGER,
    deadline INTEGER,
    exam_start VARCHAR(256),
    exam_end VARCHAR(256),
    exam_room VARCHAR(256)
);
```
Student Assignment/Project/Quiz/Exam Table, the table to store the students' submission of assignmeny/project, this table also stores the students' scores.
score_status: 1: Waiting, 2: Submitted, 3: Confirmed.
```
CREATE TABLE IF NOT EXISTS stu_assignment_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    assignment_id INTEGER NOT NULL REFERENCES assignment_exam_tab(id),
    stu_id VARCHAR(256) NOT NULL,
    type INTEGER,
    last_submit_time INTEGER,
    content TEXT,
    score INTEGER,
    score_status INTEGER
);
```
Notification Sender Table
receiver_features: a string list, each element represent a feature.
is_all_users, is_all_students, is_all_normal_students, is_all_professors, is_all_normal_professors, course_id, course_module_id, receiver_role_ids

receiver_role_ids is also a list, like ["ST20150100126", "FA19980320082"]
```
CREATE TABLE IF NOT EXISTS noti_sender_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    sender_id VARCHAR(256) NOT NULL REFERENCES role_tab(role_id),
    title VARCHAR(256),
    content TEXT,
    send_time INTEGER,
    receiver_features TEXT
);
```
Notification Receiver Table status: 1: Unread, 2: Have Read, 3: Deleted
```
CREATE TABLE IF NOT EXISTS noti_receiver_tab (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    sender_id VARCHAR(256) NOT NULL REFERENCES role_tab(role_id),
    title VARCHAR(256),
    content TEXT,
    send_time INTEGER,
    receiver_id VARCHAR(256) NOT NULL REFERENCES role_tab(role_id),
    status INTEGER
);
```
Above all are the database tables used in this project, the overall ER chart is
![ER Chart](image/system_database_er.png)

(Usually, there should be a sequence chart to illustrate the workflow, but this project is quite simple and does not involve interaction between multiple systems, the sqquence chart is not)

# API Design

## Semester Related
| API URL | Method |
|:---|:---|
| /api/v1/create_semester | POST |

To create a new semester, it will check the role permission in the request body. (**Note**: This is not a reasonable approach, but as mentioned earlier, in order to simplify the project complexity. All permission checks after this are the same.)

Request Body
```
{
    "role_id": , // a string, the role id of the semester creator
    "semester": , // a string, format follows the requirement in requirement document
    "start_time": , // an integer timestamp, must be a Monday. But the API won't check it, it should be FE to limit user to select Monday only and convert to a timestamp. It must be later than the end_time of the last semester.
    "end_time": , // an integer, for normal semester, the time duration must be within 16-20 weeks while short semester must be within 2-6 weeks
    "type": // an integer, 1: normal semester, 2: short semester, 0: vacation
}
```

Response Body
```
{
    "error_code": // an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/get_current_semester | GET |

Response Body
```
{
    "semester": , 
    "start_time": , 
    "end_time": , 
    "type": ,
    "error_code": // an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```

## Course Related
| API URL | Method |
|:---|:---|
| /api/v1/create_course | POST |

To create a course. A college couldn't create 2 courses with the same name.

Request
```
{
    "course_name": , // a string
    "college_id": , // a string
    "credit": , // an integer
    "brief": , // a string 
    "creator_role_id": // the creator id, to do permission check
}
```

Response Body
```
{
    "course_id": ,// if the course is successfully created, return the course_id, otherwise return null
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/get_course?offset=0&limit=10&college_id=001&course_id=CS40045 | POST |

`offset`, `limit`, `college_id` and `course_id` are Query String Parameters. So this API supports pagination and filtering based on college_id and course_id. *All of them can be empty. If any of them is empty, it would be treated as default value.* 

(The following APIs with Query String Parameters will only list the parameters in the table.)

Query String Parameters
| Parameter | Type | Default Value |
| :---: | :---: | :---: |
| offset | int | 0 |
| limit | int | 10 |
| college_id | string | null |
| course_id | string | null |
| course_name | string | null |

Request Body
```
{
    "role_id": //can't be empty
}
```

Response Body
```
{
    "course_list": [
        {
            "course_id": ,
            "course_name": ,
            "college_id": ,
            "credit": ,
            "brief": 
        },
        {
            "course_id": ,
            "course_name": ,
            "college_id": ,
            "credit": ,
            "brief": 
        }
    ]
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/edit_course | POST |

To edit a course. 

Request
```
{
    "course_id": , // the course id to be edited, can't be empty
    "course_name": , // a string
    "college_id": , // a string
    "credit": , // an integer
    "brief": , // a string 
    "creator_role_id": // the creator id, to do permission check
    "recommended_year": // an integer
}
```
Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```

## Course Module Related
| API URL | Method |
|:---|:---|
| /api/v1/create_course_module | POST |

To create a course module based on an existing course. 

Request Body
```
{
    "course_id": ,
    "creator_role_id": , // the role id of the creator, can't be empty
    "professor_id": ,// the professor role id of this course module, can't be empty
    "ta_id": ,// the role id of the TA of this course module, can be empty
    "semester": , //can't be empty
    "class_room": ,// string
    "class_period_start": ,// string, the API will check the format
    "class_period_end": ,// string, the API will check the format
    "duration": , //integer, 
    "course_capacity": , //integer, can't be empty
    "min_stu_num": , //integer, can't be empty
    "score_ratio": // string, can't be empty, the API will check the format, like [{"type":1, "ratio":0.1},{"type":1, "ratio":0.1}, {"type":1, "ratio":0.1}, {"type":2, "ratio":0.2}, {"type":3, "ratio":0.2}, {"type":4, "ratio":0.3}]
}
```
Response Body
```
{
    "course_module_id": // if the course module is successfully created, return the course_id, otherwise return null
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/edit_course_module | GET |

To edit an existing course module.
Request Body
```
{
    "course_module_id": , //can't be empty
    "editor_role_id": , // the role id of the creator, can't be empty
    "professor_id": ,// the professor role id of this course module, can't be empty
    "ta_id": ,// the role id of the TA of this course module, can be empty
    "semester": , //can't be empty
    "class_room": ,// string
    "class_period_start": ,// string, the API will check the format
    "class_period_end": ,// string, the API will check the format
    "duration": , //integer, 
    "course_capacity": , //integer, can't be empty
    "min_stu_num": , //integer, can't be empty
    "score_ratio": // string, can't be empty, the API will check the format, like [{"type":1, "ratio":0.1},{"type":1, "ratio":0.1}, {"type":1, "ratio":0.1}, {"type":2, "ratio":0.2}, {"type":3, "ratio":0.2}, {"type":4, "ratio":0.3}]
}
```
Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/get_course_module | GET |

Query String Parameter
| Parameter | Type | Default Value |
| :---: | :---: | :---: |
| offset | int | 0 |
| limit | int | 10 |
| college_id | string | null |
| course_id | string | null |
| course_module_id | string | null |
| professor_id | string | null |
| semester | string | {current semester} |
| class_period_start | string | null |
| class_period_end | string | null |

Request Body
```
{
    "role_id": //can't be empty
}
```

Response Body
```
{
    "course_id": ,
    "course_module_id": ,
    "professor_id": ,// the professor role id of this course module, can't be empty
    "ta_id": ,// the role id of the TA of this course module, can be empty
    "semester": , //can't be empty
    "class_room": ,// string
    "class_period_start": ,// string, the API will check the format
    "class_period_end": ,// string, the API will check the format
    "duration": , //integer, 
    "course_capacity": , //integer, can't be empty
    "min_stu_num": , //integer, can't be empty
    "score_ratio": , // string, can't be empty, the API will check the format, like [{"type":1, "ratio":0.1},{"type":1, "ratio":0.1}, {"type":1, "ratio":0.1}, {"type":2, "ratio":0.2}, {"type":3, "ratio":0.2}, {"type":4, "ratio":0.3}]
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/delete_course_module | DELETE |

to delete a course module. 

Request Body
```
{
    "course_module_id": ,// can't be empty; if the id is not existing or imvalid, the API will return error.
    "role_id": // the deleter role id, can't be empty
}
```
Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/select_course_module | POST |

to select a course module during course selection phase

Request Body
```
{
    "role_id": ,// can't be empty, this field is used for permission check
    "course_module_id": ,// can't be empty
    "stu_id": ,// can't be empty
}
```
Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/drop_course_module | POST |

to drop a selected course module

Request Body
```
{
    "role_id": ,// can't be empty, this field is used for permission check
    "course_module_id": ,// can't be empty
    "stu_id": ,// can't be empty
}
```
Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/get_selected_course_module | GET |

to get all the selected course module. A student can only get his/her selected course modules.

Query String Parameter
| Parameter | Type | Default Value |
| :---: | :---: | :---: |
| offset | int | 0 |
| limit | int | 10 |
| college_id | string | null |
| course_id | string | null |
| course_module_id | string | null |
| professor_id | string | null |
| semester | string | {current semester} |
| class_period_start | string | null |
| class_period_end | string | null |

Request Body
```
{
    "role_id": ,//can't be empty
    "stu_id": //can't be empty
}
```

Response Body
```
{
    "course_module": [
        {
            "course_id": ,
            "course_module_id": ,
            "professor_id": ,// the professor role id of this course module, can't be empty
            "ta_id": ,// the role id of the TA of this course module, can be empty
            "semester": ,
            "class_room": ,
            "class_period_start": ,
            "class_period_end": ,
            "duration": ,  
            "course_capacity": , 
            "min_stu_num": , 
            "score_ratio": , 
            "scores": ,
            "final_score": ,
            "status"
        }
    ]
    
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```

## Assignment Related
| API URL | Method |
|:---|:---|
| /api/v1/publish_assignment | POST |

to publish assignment, project, quiz, exam, using type to distinguish them.

Request Body
```
{
    "role_id": ,//can't be empty, used for permission check
    "professor_id": ,//can't be empty
    "course_module_id": ,//can't be empty
    "type": ,//can't be empty 
    "title": ,
    "content": ,
    "deadline": ,// a timestamp
    "exam_start": ,// a string, the API will check the format 
    "exam_end": ,// a string, the API will check the format 
    "exam_room": 
}
```
Response Body
```
{

    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/edit_assignment | POST |

to edit assignment, project, quiz, exam. 

```
{
    "role_id": ,//can't be empty, used for permission check
    "professor_id": ,//can't be empty
    "course_module_id": ,//can't be empty
    "type": ,//can't be empty 
    "title": ,
    "content": ,
    "deadline": ,// a timestamp
    "exam_start": ,// a string, the API will check the format 
    "exam_end": ,// a string, the API will check the format 
    "exam_room": 
}
```
Response Body
```
{
    "id": ,// the assignment id, if the assignment is published successfully, the API will return it.
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/submit_assignment | POST |

to submit an assignment or project.

Request Body
```
{
    "assignment_id": ,// can't be empty
    "role_id": ,// can't be empty, used for permission checking
    "stu_id": ,// can't be empty
    "type": ,
    "content": // can't be empty
}
```

Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/get_submitted_assignment | GET |

to get the submitted assignments

Query String Parameter
| Parameter | Type | Default Value |
| :---: | :---: | :---: |
| offset | int | 0 |
| limit | int | 10 |
| course_id | string | null |
| course_module_id | string | null |
| stu_id | string | null |
| assignment_id | int | no default value, can't be null |
| semester | string | {current semester} |

Request Body
```
{
    "role_id": ,//can't be empty
}
```

Request Body
```
{
    "submitted": [
        {
            "stu_assignment_id": ,
            "course_id": ,
            "course_module_id": ,
            "semester": ,
            "type": ,
            "last_submit_time": ,
            "content": ,
            "score": ,
            "score_status": 
        },
    ],
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/submit_score | POST |

to submit a score of an assignment/project/quiz/exam, before confirming, the score can be submitted many times

Request Body
```
{
    "role_id": ,// can't be empty
    "stu_assignment_id": ,// can't be empty
    "score": // can't be empty
}
```

Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/confirm_score | POST |

to confirm a submited score

Request Body
```
{
    "role_id": ,// can't be empty
    "stu_assignment_id": ,// can't be empty
}
```

Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
## Noti Related
| API URL | Method |
|:---|:---|
| /api/v1/send_noti | POST | 

to send a noti

Request Body
```
{
    "sender_id": ,// the role id of the sender, can't be empty
    "title": ,// if the content is empty then it can't be empty 
    "content": ,// if the title is empty then it can't be empty
    "receiver_feature": // a string, 8 items separated by semicolons: "1;0;0;0;0;0;0;0;[]"; the logical relationship between items would be viewed as "OR". For the first 5 items, 0 means no while 1 means yes. For 6th and 7th, 0 means not to use this feature while specific id means to use this feature. For the last item, [] means not to use it while valid role id list means to use this feature. 
}
```

Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/get_sent_noti | GET |

to get all the noti sent by himself/herself.

Query String Parameter
| Parameter | Type | Default Value |
| :---: | :---: | :---: |
| offset | int | 0 |
| limit | int | 10 |
| send_time_start | int | null |
| send_time_end | int | null |
| key_word | string | null | 

`send_time_start` and `send_time_end` forms a time range of send time to filter all the sent noti in this period. `key_word` is used to filter the noti whose title or content contains such key word.

Request Body
```
{
    "role_id": // can't be empty
}
```

Response Body
```
{
    "noti": [
        {
            "title": ,
            "content": ,
            "receiver_feature": ,
            "send_time":
        },
    ],
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/get_received_noti | GET |

to get all the received noti

Query String Parameter
| Parameter | Type | Default Value |
| :---: | :---: | :---: |
| offset | int | 0 |
| limit | int | 10 |
| send_time_start | int | null |
| send_time_end | int | null |
| key_word | string | null |
| status | int | null |

Request Body
```
{
    "role_id": // can't be empty
}
```

Response Body
```
{
    "noti": [
        {
            "id": ,
            "title": ,
            "content": ,
            "receiver_feature": ,
            "send_time": ,
            "status": 
        },
    ],
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/delete_noti | DELETE |

Request Body
```
{
    "role_id": // can't be empty
    "ids": [] // a list of received noti id, can't be empty or []
}
```

Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
## Phase Related
| API URL | Method |
|:---|:---|
| /api/v1/start_course_selection | POST |

to start a course selection phase.

Request Body
```
{
    "role_id": ,// can't be empty
    "end_time": // an integer timestamp; the course selection phase would be ended by the cron job after this time
}
```

Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/end_course_selection | POST |

to forcely end the course selection phase; if the current phase is not a course selection phase, it will return error.

Request Body
```
{
    "role_id": ,// can't be empty
}
```

Response Body
```
{
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
---
| API URL | Method |
|:---|:---|
| /api/v1/get_phase | GET |

To get the current phase. No permission checking.

Response Body
```
{
    "type": ,// 0 means normal phase, 1 means course selection phase
    "end_time": ,
    "error_code": ,// an integer, it would be null if no error 
    "error_message": // a string, it would be null if no error
}
```
# Global Variable
Considering that the current semester and the phase info would be frequently used for varification, singleton is an ideal data structure for them.

current semester:
```
type CurrentSemester struct{
    semesterType int
    semester string
    startTime int64
    endTime int64
}
```

phase:
```
type CurrentPhase struct{
    phaseType int
    endTime int64
}
```