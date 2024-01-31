# Authentication System
## Project Preface
The Management System used a very naive permission checking logic to determine if a request can access confidential resource. But actually the permission check should be done by a separated Authorization Service. 

For one hand, the microservice architecture is widely used in the current project, which means the authentication in each monolithic service would cause redundent development and much trivial traffic. 

For another, the non-centralized authentication logic would be a big obstacle for the unified identity management. Image that, one day your management system becomes very big. There are a lot of user data in your database and many other functions, more than teaching affairs, like booking lecture seat or book borrowing, are developed. Then the requirement "one account, access to all" is imperative. By using a separated authentication system to implement unified identity management would be an ideal solution.

Of course, at present, if you have only been exposed to Management System project, you may not have a personal experience with the above statements. This is because the Management System is too simple, it only has 1 service.

So this project won't design many APIs. Instead, the complexity of this project will mainly be reflected in the architecture. You need to implement a business service and an authentication service. The communication between services would be implemented by gRPC. Also this project will introduce a middleware, Redis.

## Overview
| Project Info | Details | Remark |
| :---- | :---- | :---- |
| Used Skilss | Golang, database, web server, authentication, cache, gRPC | |
| Difficulty | ⭐⭐⭐⭐ | A little bit higher than basic program design. |
| Recommended Completion Time | the 3rd semester of college | After basic CRUD, you will acquire more in architecture. |
| Recommended Completion Duration | 2 ~ 3 Weeks | |
| What Would You Gain? | Ability to sweep all course designs during undergraduate studies. |  |

## Note
As you have already known that before actually coding, requirement document and technology design document are needed for clarification of expected functions and specific implementation design, and the requirement of this project is very to understand, only technology design document will be provided. (Reduce redundant work)

By the way, most of the commercial company would integrate the authentication function into their gateway. But the implementation in this project is enough for you to understand how it works.