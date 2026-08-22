feat(api): Show stored greeting

Backend now serves stored greeting from PostgreSQL with documented error codes and no extra route. Frontend already fetches API data; mock removal stays done on this branch.

Verification:
- `check_build` backend ✅
