#include<bits/stdc++.h>
#include<iostream>
using namespace std;
class copyCons {
    int *a;
    public:
    copyCons() {
         cout<<"copyCons constructor called "<<endl;
    }
    copyCons(const copyCons &obj){
          cout<<"COPY called "<<endl;
        a = new int[sizeof(obj.a)];
        for(int i = 0 ; i < sizeof(obj.a) ;i++){
            a[i] = obj.a[i];
        }
    }
    void PrintA(){
        for(int i = 0 ; i < sizeof(a) ;i++){
           cout<< "value is a ["<<i<<"] is : "<<a[i] <<endl;
        }
    }
    void setA(int _size) {
        a = new int[_size];
        for(int i = 0 ; i < _size ;i++){
         a[i] = i;
        }
    }
   void operator=(const copyCons &obj){
       cout<<"operator called "<<endl;
        a = new int[sizeof(obj.a)];
        for(int i = 0 ; i < sizeof(obj.a) ;i++){
            a[i] = obj.a[i];
        }
    }
    virtual ~copyCons(){
         cout<<" copyCons Desctructor called "<<endl;
    }
};
auto fet(){
    return std::make_unique<int>();
}
class xyz :public copyCons{
   public:
    xyz() {
        cout<<"xyz constructor called "<<endl;
    }
    ~xyz(){
        cout<<"xyz Desctructor called "<<endl;
    }
};

int main()
{
    /*copyCons obj1,obj3;
    obj1.setA(5);
    obj1.PrintA();
    copyCons obj2 = obj1;
    obj2.PrintA();
obj3 = obj2;
obj3.PrintA();*/
    //copyCons *ptr = new xyz();
    //delete ptr;
    auto a =std::make_unique<int>(12);
    //a.push_back(12);
    //cout<a<<endl;
    printf("%d",*a);
    return 0;
}