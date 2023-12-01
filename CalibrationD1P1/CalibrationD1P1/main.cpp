#include "Funcs.hpp"
#include <string>
int main() {
	std::ifstream file("Input.txt");
	std::string line;
	int calibration = 0;
	while (std::getline(file, line)) {
		calibration += getCalibration(line);
	}
	std::cout << calibration;
}