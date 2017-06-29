//
// Created by wubo on 2017/6/29.
//

#ifndef CRFRAMEWORK_CRSTRINGTOKEN_H
#define CRFRAMEWORK_CRSTRINGTOKEN_H

#include <string>
#include <vector>

namespace cr
{
    /**
    * @brief 非线程安全
    */
    class crStringToken
    {
    public:
        /**
        * @brief     构造函数
        * @param str 需要分割的字符串
        * @param seq 分配的字符，如";,/"等，一次支持一个或多个分隔标识
        */
        crStringToken(const std::string& str, const char* seq);
        virtual  ~crStringToken();
    public:
       /**
       * @brief  获取分隔后的字符串列表
       * @return 字符串列表
       */
        std::vector<std::string> GetStrings();

        /**
        * @brief  获取分割后的字符串数量
        */
        size_t  GetCount();

    private:
        std::vector<std::string> m_vecStrings;
    };
}



#endif //CRFRAMEWORK_CRSTRINGTOKEN_H
