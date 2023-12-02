#include "Funcs.hpp"
#include <string>
#include <chrono>

int main() {
	const auto start{ std::chrono::steady_clock::now() };
	std::ifstream file("Input.txt");
	std::string line;
	int calibration = 0;

	while (std::getline(file, line)) {
		calibration += getCalibration(line);
	}
	std::cout << calibration;

	const auto end{ std::chrono::steady_clock::now() };
	const std::chrono::duration<double> elapsed_seconds{ end - start };
	std::cout << "\n" << elapsed_seconds;
}