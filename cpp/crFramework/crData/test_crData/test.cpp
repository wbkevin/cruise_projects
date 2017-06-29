//
// Created by wubo on 2017/6/29.
//

#include <iostream>
#include "crStringToken.h"
#include "crThread.h"

int main(int argc, char* argv[])
{
    cr::crStringToken token("12;34;56",";");

    auto vec = token.GetStrings();

    std::cout << token.GetCount()  << std::endl;

    for(const auto& iter : vec)
    {
        std::cout << iter << std::endl;
    }

   // std::cout << "hello world" << std::endl;
    return  0;
}