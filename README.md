# --------------------!!!!!!!Attention!!!!!!!--------------------
- I did all of TODO and Must have but in It Would be Great If You Have I did only Swagger and Docker image because I don't know how to develop by the Clean Architecture principles
- I used Hexagonal Architecture
- Unit test cover rate = 89.5%
- In the section that says "`Read` and `Write` the data from `JSON` file", am I understanding correctly that the data should be stored in a JSON file instead of a Database?
# ---------------------------------------------------------------
# Definition of Done
- The data must consist of the following fields
| Field  | Data Type | Notes |
| ------ | --------- | ----- |
| ID | UUID |  |
| Title | String |  |
| Description | String |  |
| Created At Date Time | Date Time with Time Zone |  | 
| Image | String |  |
| Status | String  | Accept: `IN_PROGRESS` \| `COMPLETED` |

- The TODO application can `CREATE` a task with the following requirements✔️
    - The `ID`, `Title`,`Date`, and `Status` fields must be required ✔️
    - The `ID` field must be `UUID` format ✔️
    - The `Title` field must not over `100 characters` ✔️
    - The `Date` field must be `RFC3399 with Timezone` format ✔️
    - The `Status` field must accept only `IN_PROGRESS` or `COMPLETE` status ✔️
    - The `Image` field must be `Based64 Encode` format ✔️
- The TODO application can show the `LIST` of tasks with the following requirements✔️
    - Can sort the data by `Title` or `Date` or `Status` fields ✔️
    - Can search the data by `Title` or `Description` fields ✔️
- The TODO application can `UPDATE` a task with the following requirements ✔️
    - Can update a task by `ID` field ✔️
    - Can update `Title`, `Description`, `Date`, `Image`, and `Status` fields corresponding to the requirements from the `CREATE` feature ✔️

# Application Specs for API

- **Must have**✔️
    - You can use any API frameworks ✔️
        - It would be great if you developed by `GoLang` with `GinGonic` framework ✔️
    - `Read` and `Write` the data from `JSON` file ✔️
    - Design by the `RestFul` approach ✔️
    - Write the `UnitTest` ✔️

- **It Would be Great If You Have**
    - Develop by the `Clean Architecture` principles ❌
    - Develop the `API Explorer` by `Swagger` ✔️
    - Create a `Docker Image` for API ✔️
