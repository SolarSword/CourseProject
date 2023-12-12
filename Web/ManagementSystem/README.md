# Project Preface
This is a typical web XXX management system project. As a heuristic introductory project, the xxx management system has always been assigned as a major assignment or course design project by programming courses in various universities. However, a considerable number of students with poor programming foundations will be discouraged by the difficulties encountered in such projects and lose their confidence in learning, leading to a lack of effort to study related courses seriously afterwards. Utilitarianly speaking, the skill of programming is still a pathway to a job with a relatively ideal income level, even in the current generally sluggish economic context. Not to mention that programming has become a must-have skill for high-end talents, regardless of whether you will engage in computer related industries in the future. 

Therefore, this example project aims to help students who feel overwhelmed by such large assignments, and to help you clarify the thinking methods and coping strategies of this course design project from the seemingly intimidating but not actually complex management system project. After all, it's a pity to give up learning related knowledge and skills just because of one assignment.

# Project Overall
This project aims to implement an `Academic Affairs System`, which contains a series of functions related to university teaching, such as course selection, submitting assignments, and issuing notifications... For different roles, such as students and professors, their system permissions vary. 
## Overview
| Project Info | Details | Remark |
| :---- | :---- | :---- |
| Used Skilss | Golang, database, web server | |
| Difficulty | ⭐⭐⭐ | A little bit higher than basic program design. <u>Also we won't design a very complex system for real usage.</u> |
| Recommended Completion Time | the 2nd semester of college ~  the 1st semester of sophomore year | During this period, a basic understanding of programming concepts and data structures should have been established. |
| Recommended Completion Duration | 2 ~ 3 Weeks | |
| What Would You Gain? | Practical **experience** in general web development; **Ideas and methods** for simple system design; **Confidence** in programming learning; **Assignments** that can be submitted... |  |

## Aims
To implement an `Academic Affairs System`:

- According to the differences in user roles, it has the following functions
  
  - <font color=Red>**Students**</font>:

    - **Select Course**: 1. select available courses; 2. cancel course selection before the deadline.
    - **Query Course Info**: 1. student can query the available courses during course selection period; 2. student can query enrolled courses after course selection period, including course materials, course score...; 3. simple filter function is supported.
    - **View Notifications**: view the notifications sent by system administrator or the professors whose course is selected by this student.
    - **View, Download and Submit Assignment** (Optional): 1. view the assignment published by his/her professors; 2. download the materials provided by professors if any; 3. submit the finished assignments. 4. student can query the assignment score after the professor submits the results.

  - <font color=Red>**Professors**</font>

    - **Register and Publish Course**: 1. professor can register a course; 2. professor can publish a registered course in a specific semester. 
    - **View and Public Notifications**: 1. professor can sent notifications to his/her students; 2. professor can view notifications sent by administrator and himself/herself.
    - **Public and Evaluate Assignment**: 1. professor can publish an assignment to his/her students in this semester; 2. after the assignment deadline, professor can view all the assignments and evaluate.
    - **Publish and Evaluate Quiz and Exam**: 1. professor can publish quiz or exam to his/her students; 2. professor can submit the score after the quiz or exam.

  - <font color=Red>**Administrator**</font>

    - **Publish Notifications**: 1. administrator can sent notifications to all the other users; 2. administrator can view notifications sent by himself/herself.
    - **Course Operation**: 1. administrator can help a student to select/cancel a course even if the course selection deadline has been met; 2. administrator can help a professor edit the course info even if the course has been pulished; 3. administrator can suspend a course.
    - **Correct Scores**: administrator can change the scores of students even the scores has been submited and confirmed by professors.

- General functions of the system:

    - **Calculate Scores**: 1. after the professor submits all the student scores and confirms, calculate the final scores for all the students who are enrolled by this course. 2. calculate average score, median score, highest score, lowest score, score distribution for professor.
    - **Display Semester Info**: the system can display the semestor info.
    - **Change The Course Status Automatically**: the system can change the course status at particular time: 1. course selection starts/ends; 2. course score publishes.

The above mentioned functions are not all the functions of the educational administration system in practical applications. But for a course project, it is complex enough for a student to understand how a web system works and to acquire the skills to build similar systems. 

Before you start to learn this project or code by youself, there are still some documents that you need to read or write, like `project requirement ducoment` (for this project, `project confirmation ducoment` should be fine), `technology design document` or `UI design document` (if you'd like to provide frontend interface). These different documents would help you to understand the functions you need to implement, otherwise there is a high probability that you would feel that this project is difficult to start with or that you might feel that your ideas are chaotic during actual programming. Developing these good habits will help you break down complex functions and code them down one by one. Then, just like when you were a child, you gradually completed your project like piecing together toys or building blocks.

# Notes
I will put the documents in path {project_name}/doc. Also in each folder's entry file, I will try to add enough rich information in comments to help you understand what they are doing.