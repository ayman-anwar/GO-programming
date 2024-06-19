
import pandas as pd
msg=pd.read_csv('naivetext.csv',names=['message','label'])
print('The dimensions of the dataset',msg.shape)
msg['labelnum']=msg.label.map({'pos':1,'neg':0})
X=msg.message
y=msg.labelnum
print(X)
print(y)
from sklearn.model_selection import train_test_split
xtrain,xtest,ytrain,ytest=train_test_split(X,y)
print('\n The total number of training data :',ytrain.shape)
print('\n The total number of testing data :',ytest.shape)
from sklearn.feature_extraction.text import CountVectorizer
count_vect=CountVectorizer()
xtrain_dtm=count_vect.fit_transform(xtrain)
xtest_dtm=count_vect.transform(xtest)
print('\n The words or tokens in the text document \n')
print(count_vect.get_feature_names())
df=pd.DataFrame(xtrain_dtm.toarray(),columns=count_vect.get_feature_names())
from sklearn.naive_bayes import MultinomialNB
clf=MultinomialNB().fit(xtrain_dtm,ytrain)
predicted=clf.predict(xtest_dtm)

from sklearn import metrics
print('\n Accuracy of the classifier is', metrics.accuracy_score(ytest,predicted))

print('\n Confusion matrix is\n ', metrics.confusion_matrix(ytest,predicted))

print('\n The value of recall is ', metrics.recall_score(ytest,predicted))
print('\n The value of precision is ', metrics.precision_score(ytest,predicted))
