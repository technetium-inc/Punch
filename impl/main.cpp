#include <iostream>
#include <fstream>
#include <vector>

using namespace std;

vector<string> split(string str, const char *del) {
    vector<string> data;
    string current_string = "";
    for(int x=0; x < str.length(); x++){
        if(str[x] == *del){
            data.push_back(current_string);
            current_string = "";
            continue;
        }
        current_string += str[x];
    }
    return data;
}

string remove_characters(string data, const char *character) {
    string current_string;
    for(int index=0; index<data.length(); index++){
        if(data[index] != *character){
            current_string += data[index];
        }
    }
    return current_string;
}

void run_punch_code(string data) {
    vector<string> tokens = split(data, ";");
    for(int tokenIndex=0; tokenIndex<tokens.size(); tokenIndex++){
        vector<string> current_token = split(tokens[tokenIndex], " ");
        
    }
}

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
        run_punch_code(file_content);
    }
}