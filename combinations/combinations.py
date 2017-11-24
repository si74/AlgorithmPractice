#method 1 - either include or not include each item
class getCombinationsClass:

	def __init__(self,array,k):

		#initialize empty array
		self.new_array = []
		for i in xrange(k):
			self.new_array.append(0)

		self.final = []

		self.combinationUtil(array,0,self.new_array,0,k)

	def combinationUtil(self,array,array_index,current_combo, current_combo_index,k):

		if current_combo_index == k:
			self.final.append(current_combo[:])
			return
		
		if array_index >= len(array):
			return
			
		current_combo[current_combo_index] = array[array_index]

		#if current item included
		self.combinationUtil(array,array_index+1,current_combo,current_combo_index+1,k)

		#if current item not included
		self.combinationUtil(array,array_index+1,current_combo,current_combo_index,k)

#method 2 - recursively call function - not in a class
def getCombinations(array,k):
	
	#initialize empty array
	new_array = []
	for i in xrange(k):
		new_array.append(0)

	return getCombinationsUtil(array,0,new_array,0,k)

def getCombinationsUtil(array,array_index,current_combo, current_combo_index,k):

	if current_combo_index == k:
		return [current_combo[:]]
	
	if array_index >= len(array):
		return []
		
	current_combo[current_combo_index] = array[array_index]

	#if current item included & not included
	return getCombinationsUtil(array,array_index+1,current_combo,current_combo_index+1,k) + getCombinationsUtil(array,array_index+1,current_combo,current_combo_index,k)


if __name__ == "__main__":

	result = getCombinationsClass([1,2,3],2)

	print result.final

	print getCombinations([1,2,3],2)
