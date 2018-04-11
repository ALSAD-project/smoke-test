from sklearn.cluster import Perceptron
import numpy as np
import pickle


def handle(req):
    """handle a request to the function
    Args:
        req (str): request body, expecting csv input as: 2.2, 4.5, 1\n 5.5, 2.5, 0\n ....
    """
    FILENAME = 'sl.sav'
    X = []
    Y = []

    for line in req.split('\n'):
        data = line.split(',')

        X.append([float(data[0]), float(data[1])])
        if len(data) > 2:
            Y.append(float(data[2]))
    
    try:
        clf = pickle.load(open(filename, 'rb'))
    except:
        clf = Perceptron(eta0=0.1, n_iter=40, random_state=0)    

    # Guessing flow based on label data's presence
    if len(X) == len(Y):
        clf.partial_fit(np.array(X), np.array(Y))
        pickle.dump(clf, open(filename, 'wb'))
        response = None
    else:
        Y = clf.predict(X)
    
    response = np.concatenate((X, Y.T), axis=1)
    
    # TODO: stringtify 2d array in better way
    return np.array2string(np.array(response), separator=',').replace('[', '').replace('],', '').replace(']','')
