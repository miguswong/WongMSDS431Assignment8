# MSDS431 - Assignment 8: Modern Applied Statistics with Go

## Assignment Details

### Management Problem

A research consultancy specializes has been using R as its major software platform for many years, relying on the Comprehensive R Archive NetworkLinks to an external site. as a source of modeling techniques, including many computer-intensive techniques. The firm has heard about the performance advantages of Go and its ability to make full use of modern multi-core processors. Concerned about growing cloud computing costs when conducting its research (especially when working with large data sets), the firm is considering moving much of its programming work from R to Go. 

The consultancy is looking for an independent data scientist to evaluate the possibilities of Go for the firm. How easy is it to use Go in place of R? How much money will the firm save in cloud computing costs?

#### Assignment Requirements 

Evaluate the possibilities for using Go in modern applied statistics: 

* Read the Gelman and Vehtari (2021) article and select one of the computer intensive statistical methods listed in that article: counterfactual causal inference, bootstrapping and simulation-based inference, overparameterized models, regularization, Bayesian multilevel models, generic algorithms, or adaptive decision analysis. It should be a statistical method that has been fully implemented in the R programming language. [For example, consider implementing bootstrap sampling, providing a robust, distribution free, estimate of the standard error of the median for distributions of various shapes (positively skewed, symmetric, and negatively skewed). Start by implementing in R and then refactor into Go. See the R example at the end of this write-up. Note: if you work on a bootstrap estimator for the median, you may be able to use some of your code in the Week 9 programming assignment.]
* Referring to the [Comprehensive R Archive Network](https://cran.r-project.org/), identify the R package(s) for the selected method and demonstrate an analysis in R. Name the method and show its web addresses (URLs).
* Implement the selected statistical method in Go either by drawing on a third-party Go package or developing your own Go program.  Sites that may be of special value in responding to this discussion question are [Awesome Go](https://awesome-go.com/gui/), and [go.dev](https://pkg.go.dev).
* Employ Go testing, benchmarking, software profiling, and logging. What did you do to improve the performance of the Go implementation of the selected statistical method?
* Compare the R and final Go implementations using the same input data. Ensure that results are comparable. Report on the memory requirements and processing times associated with the R and final Go implementations.
In the **README.md** file of the repository, describe your efforts in finding R and Go packages for the method. Review your process of building the Go implementation. Review your experiences with testing, benchmarking, software profiling, and logging.
* Finally, in the **READMe.md** file, make a recommendation to the research consultancy. Under what circumstances would it make sense for the firm to use Go in place of R for the selected statistical method? Select a cloud provider of infrastructure as a service (IaS). Note the cloud costs for virtual machine (compute engine) services. What percentage of cloud computing costs might be saved with a move from R to Go?

## Chosen Statistical Method: Bootstrap Method
The bootstrap estimator for median was  the chosen statistical method and implementd in both R and Go.

The implementation of R was heavily based on Dr. Miller's [run-bootstrap-median.R](https://github.com/miguswong/WongMSDS431Assignment8/blob/main/run-bootstrap-median.R). While this program only utilizes R's standard packages, it should be noted that there are other pre-developed bootstrap packages such as [Efron and Tibshirani's](https://cran.r-project.org/web/packages/bootstrap/index.html) or the [boot package](https://cran.r-project.org/web/packages/boot/index.html)

The [go package](https://github.com/miguswong/WongMSDS431Assignment8/blob/main/main.go) developed was heavily based off the R script and used to achieve similar results; the output of both programs are in the same format. However, there are a couple of things to note in terms of differences between the programs.
* R and Go do not use the same random seed generator therefore, results from the program are not directly comparable.
    * Because of this and the fact that rnorm generates seeds differently from how random normal numbers are generated in Go create different results.
* Many mathematical operations such as Mean, Median, and Standard Deviation are not readily available within R. So, they were manually created.

## Development of Go Program
Several packages were utilized in the development of the program.

````go
import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
    

	"github.com/seehuhn/mt19937"
)
````

While many are typically used in many projects, it is important to touch on *math* and [*github.com/seehuhn/mt19937*](https://github.com/seehuhn/mt19937). 

Go's standard library actually does contain a math package which was extensively used across the application for arithmetic calculations. This included random number generation and simple exponential functions (square roots).

[*github.com/seehuhn/mt19937*](https://github.com/seehuhn/mt19937) was utilized in an attempt to exactly match how R generates Random numbers by default which is done with the Mersenne Twister RNG method. However, results proved that the way random numbers are generated in Go differ from R's *rnorm* function thus causing an inconsistency between the two applications.

## Comparison of Performance Results
During testing, it was found that Go was able to significantly execute faster compared to R. While Go typically executed the application in ~0.9s, R typically took magnitudes longer, around 17 s for similar results

Go was also able to mroe efficiently utilize memory. The total Allocated memory for the program as 122 MiB whereas R was utilizing nearly 211 MiB. This stark difference could be contributed to Go's built in garbage collector which helps efficiently manage memory consumption and usage.

## Discussion of Go Development
While the performance of Go was significantly faster in terms of execution when compared to R. The ease of implementation was nowhere near as straightforward. While R has many packages already available for statistical analysis, Go does not have the same level of community or support for these types of packges. Because of this, development can be seen as somewhat manual. Another interesting note that may hinder this type of development is static typing in go. Because Go requires variable to be predefined before compilation (with the exception of utilizing gnenerics), writing functions can become somewhat cumbersome. Another aspect that should not be overlooked is the functionality of dataframes which are common in both Python and R. While I was able to utiliize slices for iterating throught the generated data, it was nowhere near as straightforward compared to the easy manipulation that can be done in Python/R.

## Reccomendation to Leadership
It may make sense to utilize Go for bootstrapping over R in scenarios in which *speed and efficiency* are imperative such as real-time data analytics or any other time-sensitive scenario. Furthermore, if a lean/efficeint process is needed for cloud computing, Go will be able to help cut down on costs of data processing. What you get for in terms of efficeny and speed, you lose in ease-of-use and overall breadth of support. If Go is chosen for developing statistical calculations, the firm should be prepared to create cutom packages that may not be readily available.