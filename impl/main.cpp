#include <iostream>
#include <fstream>

using namespace std;


string find_filename(string filename) {
    string file_text;
    string content;
    int line_count = 0;
    ifstream punch_source(filename);

    while (getline(punch_source, file_text)) {
        line_count += 1;
        content += file_text;
    }

    punch_source.close();
    return content;
}

void throw_punch_error(string message, string type, int line_number) {
    cout << "Found an Error:" << type << endl;
    cout << message << "at line" << line_number;
}

int main(int argc, char *argv[]){
    if (argc != 2) {
        throw_punch_error(
            "Missing required parameter - filename",
            "ArgumentError",
            0
        );
    } else {
        string file_content = find_filename(argv[1]);
        cout << file_content;
    }
}