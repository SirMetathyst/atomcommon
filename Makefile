prog = atom component add -p atomcommon

generate:
	$(prog) -n Position2 -d x:float32 -d y:float32 
	$(prog) -n Velocity2 -d x:float32 -d y:float32 