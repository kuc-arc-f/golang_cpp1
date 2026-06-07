#include <iostream>
#include <string>
#include "http_client.hpp"

class MyTest {
private:
    std::string m_name;
public:
    explicit MyTest(std::string str){}
    ~MyTest() {}

    static void print_response(const std::string& label, const HttpResponse& resp)
    {
        std::cout << "\n===== " << label << " =====\n";
        if (!resp.error.empty()) {
            std::cerr << "[ERROR] " << resp.error << "\n";
            return;
        }
        std::cout << "Status : " << resp.status_code << "\n";
        std::cout << "Body   :\n" << resp.body << "\n";
    }

    void test() {
        HttpClient client(30 /*timeout*/, true /*verify_ssl*/);
        try{
            std::string url = "http://localhost";
            auto resp = client.get(url);
            print_response("GET-BODY:", resp);
        } catch (const std::exception& e) {
            std::cout << "Error , main" << std::endl;
            return;
        }        
      //std::cout << str << " \n";
    }


};

