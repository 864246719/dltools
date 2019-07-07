
import setuptools
with open("README.md", "r") as fh:
    long_description = fh.read()
setuptools.setup(
     name='dltools',  
     version='0.1',
     scripts=['dltools'] ,
     author="guhai",
     author_email="bysideen@gmail.com",
     description="common tools",
     long_description=long_description,
   long_description_content_type="text/markdown",
     url="https://github.com/GuNanHai/dltools_pkg",
     packages=setuptools.find_packages(),
     classifiers=[
         "Programming Language :: Python :: 3",
         "License :: OSI Approved :: MIT License",
         "Operating System :: OS Independent",
     ],
 )
