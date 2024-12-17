from fastapi import FastAPI, HTTPException, Depends
from pydantic import BaseModel
import psycopg2
from psycopg2.extras import RealDictCursor
from datetime import datetime
from typing import List, Dict

# FastAPI app
app = FastAPI()

# Dependency to get DB connection
def get_db():
    try:
        conn = psycopg2.connect(**DATABASE_CONFIG, cursor_factory=RealDictCursor)
        yield conn
    finally:
        conn.close()

# Pydantic models
class RationCreate(BaseModel):
    user_id: int
    ration: str

# Endpoints
@app.post("/ration", response_model=dict)
def create_ration(ration: RationCreate, db=Depends(get_db)):
    with db.cursor() as cursor:
        query = """
        INSERT INTO daily_ration (user_id, ration, created_at)
        VALUES (%s, %s, %s) RETURNING id;
        """
        cursor.execute(query, (ration.user_id, ration.ration, str(datetime.utcnow())))
        db.commit()
        ration_id = cursor.fetchone()["id"]
        return {"message": "Ration entry created successfully", "id": ration_id}

@app.delete("/ration/{ration_id}", response_model=dict)
def delete_ration(ration_id: int, db=Depends(get_db)):
    with db.cursor() as cursor:
        query = "DELETE FROM daily_ration WHERE id = %s RETURNING id;"
        cursor.execute(query, (ration_id,))
        deleted_id = cursor.fetchone()
        if not deleted_id:
            raise HTTPException(status_code=404, detail="Ration entry not found")
        db.commit()
        return {"message": "Ration entry deleted successfully"}

@app.get("/ration/{user_id}", response_model=List[Dict])
def get_rations(user_id: int, db=Depends(get_db)):
    with db.cursor() as cursor:
        query = "SELECT id, ration, created_at FROM daily_ration WHERE user_id = %s;"
        cursor.execute(query, (user_id,))
        rations = cursor.fetchall()
        if not rations:
            raise HTTPException(status_code=404, detail="No rations found for this user")
        return rations

class UserCharacteristics(BaseModel):
    user_id: int
    upper_strength: float
    lower_strength: float
    flexibility: float
    endurance: float
    height: float
    weight: float
    imt: float

@app.post("/user/characteristics", response_model=dict)
def create_user_characteristics(user_characteristics: UserCharacteristics, db=Depends(get_db)):
    with db.cursor() as cursor:
        query = """
        INSERT INTO user_characteristics (user_id, upper_strength, lower_strength, flexibility, endurance, height, weight, imt, created_at)
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s) RETURNING id;
        """
        cursor.execute(query, (user_characteristics.user_id, user_characteristics.upper_strength, user_characteristics.lower_strength, user_characteristics.flexibility, user_characteristics.endurance, user_characteristics.height, user_characteristics.weight, user_characteristics.imt, str(datetime.utcnow())))
        db.commit()
        user_characteristics_id = cursor.fetchone()["id"]
        return {"message": "User_characteristics entry created successfully", "id": user_characteristics_id}

@app.get("/user/characteristics/{user_id}", response_model=List[Dict])
def get_user_characteristics(user_id: int, db=Depends(get_db)):
    with db.cursor() as cursor:
        query = "SELECT id, user_id, upper_strength, lower_strength, flexibility, endurance, height, weight, imt, created_at FROM user_characteristics WHERE user_id = %s ORDER BY created_at DESC;"
        cursor.execute(query, (user_id,))
        user_characteristics = cursor.fetchall()
        if not user_characteristics:
            raise HTTPException(status_code=404, detail="No rations found for this user")
        return user_characteristics

class Equipment(BaseModel):
    name: str

@app.get("/workout/equipment", response_model=List[Dict])
def get_all_equipment(db=Depends(get_db)):
    with db.cursor() as cursor:
        query = "SELECT id, name FROM equipment;"
        cursor.execute(query)
        equipment = cursor.fetchall()
        if not equipment:
            raise HTTPException(status_code=404, detail="No rations found for this user")
        return equipment

class Workout(BaseModel):
    workout_name: str
    user_id: int
    upper_strength_goal: float
    lower_strength_goal: float
    flexibility_goal: float
    endurance_goal: float
    kcal_restriction: int