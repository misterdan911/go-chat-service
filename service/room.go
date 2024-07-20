package service

func GetRoomList(limitParam int) error {
	/*
		roomCollection := db.DB.Collection("rooms")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Extract limit from query parameters
		limit := int(30)
		if limitParam != 0 {
			limit = limitParam
		}

		// Define the filter
		filter := bson.M{
			"people": bson.M{"$in": bson.A{user.ID}},
			"$or": bson.A{
				bson.M{"lastMessage": bson.M{"$ne": nil}},
				bson.M{"isGroup": true},
			},
		}

		// Define the options
		opts := options.Find()
		opts.SetSort(bson.M{"lastUpdate": -1})
		opts.SetLimit(limit)
		opts.SetProjection(bson.M{
			"people.email":    0,
			"people.password": 0,
			"people.friends":  0,
			"people.__v":      0,
		})

		// Find the rooms
		cursor, err := collection.Find(ctx, filter, opts)
		if err != nil {
			http.Error(w, `{"error": true}`, http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)

		// Prepare the results
		var rooms []Room
		if err = cursor.All(ctx, &rooms); err != nil {
			http.Error(w, `{"error": true}`, http.StatusInternalServerError)
			return
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"limit": limit,
			"rooms": rooms,
		}
		json.NewEncoder(w).Encode(response)
	*/

	return nil
}
