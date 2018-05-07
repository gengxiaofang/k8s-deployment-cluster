## Git 简单使用总结
#### 看到比较喜欢的项目，想保存到自己本地，这时候就 clone 一下
```
git clone https://github.com/XiaoMuYi/ops_doc.git  ## 换成你想clone的地址；
```
#### 然后修修改改就是自己的，这时候我们想长久保存
```
git add ops_doc               ## 就是指定自己想提交的内容；  
git commit -m "update files"  ## 备注一下自己提交了什么内容；  
git push                      ## 上传到自己的git或github仓库；  
```
#### 上传一般需要验证，所以要配置自己的用户名或者邮箱地址
```
git config --global user.name "XiaoMuYi"
git config --global user.email "yang972711021@gmail.com"
```
