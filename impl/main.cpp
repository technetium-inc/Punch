#include <iostream>
#include <fstream>
#include <vector>

using namespace std;

/**
 * Evokes a punch exception
 * 
 * @param message The error message
 * @param type The type of error . Runtime, Argument etc.
 * @param line_number The line at which the error occured
 * 
 * @returns void
*/
void throw_punch_error(string message, string type, int line_number)
{
    cout << "Found an Error:" << type << endl;
    cout << message << "at line" << line_number;
}

/**
 * Split a string by a delimiter
 * 
 * @param str The string to split
 * @param del The del to split the string by
 */
vector<string> split(string str, const char *del)
{
    vector<string> data;
    string current_string = "";
    for (int x = 0; x < str.length(); x++)
    {
        if (str[x] == *del)
        {
            data.push_back(current_string);
            current_string = "";
            continue;
        }
        current_string += str[x];
    }
    return data;
}

string remove_characters(string data, const char *character)
{
    string current_string;
    for (int index = 0; index < data.length(); index++)
    {
        if (data[index] != *character)
        {
            current_string += data[index];
        }
    }
    return current_string;
}

void print(string print_text)
{
    cout << print_text << endl;
}

bool run_punch_code(string data)
{
    vector<string> tokens = split(data, ";");
    int lives = 100;
    for (int tokenIndex = 0; tokenIndex < tokens.size(); tokenIndex++)
    {
        vector<string> current_token = split(tokens[tokenIndex], " ");
        bool print_text = true;
        if (current_token.size() == 2)
        {
            throw_punch_error("Found 0 commands", "RuntimeError", tokenIndex + 1);
            return false;
        }

        string character = current_token[0];
        if (character == "TrashTalkLn[str]")
        {
            lives -= 5;
        }
        else if (character == "TrashTalk[str]")
        {
            lives -= 5;
        }
        else if (character == "TrashTalk[int]" || character == "TrashTalkLn[int]")
        {
            lives -= 5;
            print_text = false;
            print(current_token[1]);
        }
        else if (character == "FaceSlap[var,int]")
        {
            lives -= 4;
        }

        if (lives > 0 && print_text)
        {
            print(remove_characters(
                remove_characters(current_token[1], "\""),
                "\\"));
        }
    }
    return true;
}

string find_filename(string filename)
{
    string file_text;
    string content;
    int line_count = 0;
    ifstream punch_source(filename);

    while (getline(punch_source, file_text))
    {
        line_count += 1;
        content += file_text;
    }

    punch_source.close();
    return content;
}

int main(int argc, char *argv[])
{
    if (argc != 2)
    {
        throw_punch_error(
            "Missing required parameter - filename",
            "ArgumentError",
            0);
    }
    else
    {
        string file_content = find_filename(argv[1]);
        run_punch_code(file_content);
    }
}