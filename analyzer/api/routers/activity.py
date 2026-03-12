from fastapi import APIRouter

router = APIRouter(tags=["activity"])

@router.get("/getuser")
def get_activity_for_user():
    print("get activity")
    return {"msg": "get activity"}
