db.lands.find({ "crop_stage.crop_type": "Lúa mùa" })

db.land_objects.find({ "land_id": ObjectId("xxx") })

db.lands.find({
  "location": {
    "$geoIntersects": {
      "$geometry": { "type": "Point", "coordinates": [105.86, 21.02] }
    }
  }
})
