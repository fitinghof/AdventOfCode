#include "Funcs.hpp"

int getCalibration(std::string line) {
	int firstNum;
	int secondNum;

	std::string numbers[9]{ "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };

	for (int i = 0; i < line.size(); i++) {
		if (line[i] >= '0' && line[i] <= '9') {
			firstNum = line[i] - '0';
			break;
		}

		bool numFound = false;
		for (int k = 0; k < 9; k++) {

			if(i + numbers[k].size() < line.size()) {
				for (int j = 0; j < numbers[k].size(); j++)
				{
					if (line[i + j] != numbers[k][j]) {
						break;
					}
					if (j == numbers[k].size() - 1) {
						numFound = true;
						firstNum = k + 1;
					}
				}
				if (numFound) break;
			}
		}
		if (numFound) break;
	}
	for (int i = line.size()-1; i >= 0; i--) {
		if (line[i] >= '0' && line[i] <= '9') {
			secondNum = line[i] - '0';
			break;
		}

		bool numFound = false;
		for (int k = 0; k < 9; k++) {
			if (i - numbers[k].size() > 0) {

				for (int j = 0; j < numbers[k].size(); j++)
				{
					if (line[i - j] != numbers[k][numbers[k].size() - 1 - j])
						break;
					if (j == numbers[k].size() - 1) {
						numFound = true;
						secondNum = k + 1;
					}
				}
				if (numFound) break;
			}
		}
		if (numFound) break;
			
	}
	return firstNum * 10 + secondNum;
}
