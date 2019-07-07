#!/usr/bin/env python
# -*- coding:utf-8 -*-
# Author: wxnacy(wxnacy@gmail.com)
# Description:

from setuptools import setup, find_packages

setup(
    name = 'dlpytools',
    version = '0.0.1',
    keywords='tools',
    description = 'common tools',
    license = 'MIT License',
    url = 'https://github.com/GuNanHai/dltools',
    author = 'guhai',
    author_email = 'bysideen@gmail.com',
    packages = find_packages(),
    include_package_data = True,
    platforms = 'any',
    install_requires = [
        'requests>=2.19.1',
        'pycrypto>=2.6.1',
        'xmltodict>=0.11.0'
        ],
)
