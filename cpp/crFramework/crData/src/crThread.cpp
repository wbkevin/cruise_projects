//
// Created by wubo on 2017/6/29.
//

#include "crThread.h"


namespace cr
{
    crThread::crThread()
    {

    }

    crThread::~crThread()
    {

    }

    void crThread::Start()
    {

    }

    void crThread::Stop()
    {

    }

    void crThread::Wait()
    {
        m_thread.join();
    }

}