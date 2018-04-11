from sklearn.cluster import Birch
import numpy as np
import pickle


def handle(req):
    """handle a request to the function
    Args:
        req (str): request body, expecting csv input as: 2, 4\n 5, 2\n ....
    """
    FILENAME = 'usl.sav'
    X = []

    for line in req.split('\n'):
        data = line.split(',')
        X.append([float(data[0]), float(data[1])])
    
    try:
        brc = pickle.load(open(filename, 'rb'))
    except:
        brc = Birch(branching_factor=50, n_clusters=2, threshold=0.5, compute_labels=True)

    brc.partial_fit(np.array(X))
    pickle.dump(brc, open(filename, 'wb'))
    
    # Predicting regardless of flow
    Y = brc.predict(X)
    response = np.concatenate((X, Y.T), axis=1)
    
    # TODO: stringtify 2d array in better way
    return np.array2string(np.array(response), separator=',').replace('[', '').replace('],', '').replace(']','')
