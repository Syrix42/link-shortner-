# Project Collaboration Rules & Development Agreements

**Project:** Link-shortner Project  
**Date:** 2026-02-24  
**Version:** 0.1 (Draft)  
**Prepared by:** Alireza maleki

## Responsibilities of Developers

Each developer will be given a task (API End point) and his responsibility is to implement it end to end.

Developer must also write the automated tests:
- Unit tests for the logic of the app to follow Assertion
- Integration Tests for APIs

And write documents if specific specification was worthed sharing with other devs.

---

## Shared points among devs in the project

Devs should be caution about the project when they see points of conflictions.

Points that codes of different devs can collide are:

- Migration files (if any update on Migrations should happen for development it must be discussed between devs and then committed)
- Main.go file of the project where dependency injection plus routing happens

Devs should be caution to have the latest commit from main branch pulled when working there because they might commit something that goes into the main branch and that commit just deletes someone Elses codes (even though commits to main will be happen from PR requests but its just better to follow this rule).

Developers must start a branch for each task they are given to implement and after finishing up those they must write a PR request and after reviewing the branch's last commit will be added to the main branch.

## Milestones

### Project Docs & Agreements
- collaboration rules (this doc)
- API conventions
- shared points rules
- branching / PR rules

### Architecture Decisions (ADRs)
- key technical decisions finalized and documented
- e.g. router, DB access pattern, migration tool, response format, error handling style

### API Implementation
- devs implement assigned endpoints end-to-end

### Testing & Bug Fixing
- unit tests
- integration tests for key APIs
- bug fixing / regression tests
- integration issues in shared points resolved

### Delivery / Finalization
- final review
- all agreed endpoints merged
- docs updated
- handoff / deployment-ready backend package
