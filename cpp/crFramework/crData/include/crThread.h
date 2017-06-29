//
// Created by wubo on 2017/6/29.
//

#ifndef PROJECT_CRTHREAD_H
#define PROJECT_CRTHREAD_H

#include <thread>

namespace cr
{

/**
* @brief 线程封装
*/

    class crThread
    {
    public:
        crThread();
        virtual ~crThread();
    public:
        /**
        * @brief 线程启动接口
        */
        void Start();

        /**
        * @brief 线程结束接口
        */
        void Stop();

        /**
        * @brief 等待线程结束
        */
        void Wait();

    protected:
        /**
        * @brief   线程运行函数
        * @return  true -> 成功 false ->失败（退出）
        */
        virtual bool Run() = 0;

    private:
        std::thread m_thread;
    };
}


#endif //PROJECT_CRTHREAD_H
