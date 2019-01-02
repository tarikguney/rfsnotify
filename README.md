# R-FSNotify

[![Build Status](https://travis-ci.org/tarikguney/rfsnotify.svg?branch=master)](https://travis-ci.org/tarikguney/rfsnotify)

> This project is still under development and its API is not stable. Please use it with caution or wait until this message is removed and the first version is released.

A recursive file watcher package based on `github.com/fsnotify/fsnotify`.

Unfortunately, `fsnotify` does not have recursive watching capability, and you need to write your way to find all the sub-folders and files underneath and add them by using its `Add()` method.

R-FsNotify is the solution for that problem. It automatically watches all of your files under a directory.

Since this package is still under development, the API surface may change as the new requirements come up. Therefore, until the first release is fully published, use this library with caution.

## Unit Tests
This project is covered by various unit tests in the `rfsnotify_test.go` file. My intention is to keep it 100% covered. In case you would like to create PR for this project, please make sure that your code does not reduce the test coverage score. However, I am aware of the fact that not everything can be unit-testable, but it is still a crucial practice to keep unit tests in mind while contributing to this project.

![logo](recursive-fsnotify-thumbnail.png)

## Note
This project is still under development and if you like you can watch the all of the live streams during which this package has been actively developed.

Watch the development series here: https://www.youtube.com/watch?v=_bePpkKfU5s&list=PL_Z0TaFYSF3J8bCnGTIdZ_YgSH2THmzZt

Developed by @tarikguney with <3