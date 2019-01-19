---
categories: null
comments: true
date: 2017-05-06T08:13:00+08:00
title: 機器學習筆記(week1)
url: /2017/05/06/machine-learning-note-week1/
---

最近朋友貼了這個課程給我:[机器学习 - 斯坦福大学](https://www.coursera.org/learn/machine-learning/home/welcome)

總共有 11 週的課程

上班前看一部分，下班後再複習做做筆記

因為裡面的理論會用到微積分，只好在休息時間問問對數學比較拿手的朋友裡面的公式是如何推導的

有一種再補大學學債的感覺

第一週的主題為：

1. Introduction
1. Linear Regression with One Variable
1. Linear Algebra Review

<!--more-->

## Introduction

### 機器學習的定義

> Tom Mitchell (1998) Well-posed Learning Problem: A computer program is said to _learn_ from experience E with respect to some task T and some performance measure P, if its performance on T, as measured by P, improves with experience E.

T: task
P: performance
E: experience

以讓程式學習如何過濾垃圾信件為例：

T: 將信件分類正常信件與垃圾信件
P: 判斷的正確率
E: 觀察你將信件標記為正常信件或垃圾信件

### 演算法

先提到有：
* 監督式學習 (Supervised learning)
	* 先前可以得知訓練資料中哪些為正確答案
	* 分類問題(Classification Problem): Discrete valued output(0 or 1)
  * 回歸問題(Regression Problem): Goal is predict a continuous valued output
* 非監督式學習 (Unsupervised learning)
	* 先前無法得知訓練資料中哪些為正確答案

## Linear Regression with One Variable

線性回歸函數：

```mathjax
h_{\theta_0}(x) = \theta_0 + \theta_1x \\
```

要如何用大量訓練資料找出適合的 

```mathjax
\theta_0, \theta_1 \\
```

使函數最接近訓練資料的分佈，進而提升函數預測的精準度

以下是 `cost function`，值越小，函數也就越接近訓練資料的分佈

```mathjax
J(\theta_0, \theta_1) = \frac1{2m}\sum_{i=1}^m(h_{\theta_0}(x^{(i)}) - y^{(i)})^2 \\
```

可以使用梯度下降法(Gradient Descent)來最小化，不過無法保證得到最佳解

```mathjax
J(\theta_0, \theta_1) \\
```

梯度下降法(Gradient Descent)的公式為：

```mathjax
	\theta_j := \theta_j - \alpha {\frac\partial{\partial\theta_i}} J(\theta_0, \theta_1) \\
  \text (for\ j = 0\ and\ j = 1) \\
  \text (\alpha\ is\ learning\ rate\ 會影響收斂的速度) \\
```

將 `cost funciton` 代入：

```mathjax
\newcommand{\hf}{h_{\theta_0}(x^{(i)}) - y^{(i)}}
\newcommand{\ptz}{\frac\partial{\partial\theta_0}}
\begin{equation}
    \eqalign{
		\ptz J(\theta_0, \theta_1) & = \ptz \frac1{2m}\sum_{i=1}^m(\hf)^2 \\
		& = 2 \times \frac1{2m}\sum_{i=1}^m(\hf) \times \ptz (\hf) \\
		& = \frac1{m}\sum_{i=1}^m(\hf) \times \ptz (\hf) \\
		& = \frac1{m}\sum_{i=1}^m(\hf) \times \ptz (\theta_0 + \theta_1x^{(i)} - y^{(i)}) \\
		& = \frac1{m}\sum_{i=1}^m(\hf) \times 1 \\
		& = \frac1{m}\sum_{i=1}^m(\hf)
    }
\end{equation}
\\
```

```mathjax
\newcommand{\hf}{h_{\theta_0}(x^{(i)}) - y^{(i)}}
\newcommand{\pto}{\frac\partial{\partial\theta_1}}
\begin{equation}
    \eqalign{
		\pto J(\theta_0, \theta_1) & = \pto \frac1{2m}\sum_{i=1}^m(\hf)^2 \\
		& = 2 \times \frac1{2m}\sum_{i=1}^m(\hf) \times \pto (\hf) \\
		& = \frac1{m}\sum_{i=1}^m(\hf) \times \pto (\hf) \\
		& = \frac1{m}\sum_{i=1}^m(\hf) \times \pto (\theta_0 + \theta_1x^{(i)} - y^{(i)}) \\
		& = \frac1{m}\sum_{i=1}^m(\hf) \times x^{(i)}
    }
\end{equation}
\\
```

最後的演算法：

```mathjax
\newcommand{\hf}{h_{\theta_0}(x^{(i)}) - y^{(i)}}
\newcommand{\pto}{\frac\partial{\partial\theta_1}}
repest\ until\ convergence\ \{ \\
\begin{equation}
    \eqalign{
    	\theta_0 &:= \theta_0 - \alpha \frac1{m}\sum_{i=1}^m(\hf) \\
      \theta_1 &:= \theta_1 - \alpha \frac1{m}\sum_{i=1}^m(\hf) \times x^{(i)}
    }
\end{equation} \\
\}
```

## Linear Algebra Review

介紹矩陣 Matrix 的基本運算

`row` 的值是由上往下遞增

`column` 的值是由左往右遞增

本門課 Matrix index 的值是由 `1` 開始遞增，而不是如程式語言由 `0` 開始遞增

標記方法為:

```mathjax
\eqalign{
&R^{m \times n} \\
&m:\ rows \\
&n:\ columns \\ 
} \\
```

Vector 是 `m x 1` 的 Matrix

```mathjax
\eqalign{
&R^{m} \\
&m:\ rows \\
} \\
```

取值的時候：

```mathjax
\eqalign{
&A_{ij} \\
&i:\ row \\
&j:\ column \\
} \\
```

Iedentity Matrix(用 `I` 表示):

一定是 

```mathjax
\eqalign{
& I\ 或\ I^{m \times m} \\
& 而且對於任何\ Matrix\ A:\\
& A \times I = I \times A = A
} \\
```

例如：

```mathjax
\begin{bmatrix}
1 & 0 & 0 \\
0 & 1 & 0 \\
0 & 0 & 1 \\
\end{bmatrix}
\\
```

Matrix inverse(If A is an `m x m` matrix, and if it has an inverse):

```mathjax
AA^{-1} = A^{-1}A = I
\\
```

Matrix Trunspose:

```mathjax
A =
\begin{bmatrix}
1 & 2 & 0 \\
3 & 5 & 9 \\
\end{bmatrix}
\\
B =
A^T =
\begin{bmatrix}
1 & 2 & 0 \\
3 & 5 & 9 \\
\end{bmatrix}
\\
B_{ij} = A_{ji} \\
```

## 參考

[Machine Learning學習日記 — Coursera篇 (Week 1.1):Supervised learning,Unsupervised learning,regression,classification](https://medium.com/@ken90242/machine-learning%E5%AD%B8%E7%BF%92%E6%97%A5%E8%A8%98-coursera%E7%AF%87-introduction-supervised-learning-unsupervised-learning-18000f2b9ca1)

[[學習筆記] 機器學習: 單變數線性回歸 (Machine learning: Linear Regression with One Variable)](http://murphymind.blogspot.tw/2017/03/linear.regression.html)

[臺大開放式課程 微積分 第12章 - 偏導數](http://ocw.aca.ntu.edu.tw/ntu-ocw/ocw/cou/100S111/42)

[在 logdown 裡面打複雜方程式](http://blog.ponan.li/post/2013/08/03/write-complex-latex-equations-in-logdown-by-mathjax-support)