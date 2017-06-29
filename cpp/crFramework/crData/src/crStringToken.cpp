//
// Created by wubo on 2017/6/29.
//

#include "crStringToken.h"
#include <string.h>


namespace cr
{
    crStringToken::crStringToken(const std::string& str, const char *seq)
    {
        if (seq != nullptr)
        {
            char* p = strtok((char*)str.c_str(), seq);
            while (p != nullptr)
            {
                m_vecStrings.push_back(p);
                p = strtok(nullptr, seq);
            }
        }
    }

    crStringToken::~crStringToken()
    {

    }

    std::vector<std::string> crStringToken::GetStrings()
    {
        return m_vecStrings;
    }

    size_t  crStringToken::GetCount()
    {
        return  m_vecStrings.size();
    }

}