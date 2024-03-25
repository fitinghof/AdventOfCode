#include "Funcs.hpp"

int getCalibration(std::string line) {
	int firstNum;
	int secondNum;

	for (int i = 0; i < line.size(); i++) {
		if (line[i] >= '0' && line[i] <= '9') {
			firstNum = line[i] - '0';
			break;
		}
	}
	for (int i = line.size()-1; i >= 0; i--) {
		if (line[i] >= '0' && line[i] <= '9') {
			secondNum = line[i] - '0';
			break;
		}
	}
	return firstNum * 10 + secondNum;
}
