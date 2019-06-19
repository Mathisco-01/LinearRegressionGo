# Linear Regression in Go
Not really built for any specific purpose, just more of a test to see if I could do it.
Turns out, I can and it's not that hard.
## Formulas
`func  cost` and `func individualCost`both return the MSE (mean square error), only difference is that `func individualCost` has an m (size) of 1 so it basically just returns the squared distance to the fit line divided by 2.

![img](https://i.imgur.com/6Qu5jjv.png)


For the gradient descent I didn't want to use derivates to calculate the slope of the cost function. So im using this formula to get a more accurate theta on each iteration. Due to the nature of not using derivates it's impossible to accurately predict the minimum so in this case it's stopped right after the cost goes back up. This shouldn't cause too much issues since it only takes one feature.

![img](https://cdn-images-1.medium.com/max/1181/1*8Omixzi4P2mnqdsPwIR1GQ.png)

## Accuracy
After some tweaking I found that it was relatively accurate. I compared it to a repository I built a couple days ago based on Andrew Ng's algorithms. Run on the same single variate dataset here are the results:

This project: `Current theta:  [0.08975785196462276 0.7646112377458362]  cost:  5.896308082090673`

Andrew Ng: `Current theta:  [0.052211 0.735307]  cost:  6.052326`

With some tweaking my algorithm and hyperparameters could probably be improved to further match Andrew Ng's model. So all and all I would call this experiment a huge success!