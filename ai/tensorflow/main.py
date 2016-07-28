# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

from __future__ import absolute_import
from __future__ import division
from __future__ import print_function

import tempfile
import pandas as pd
import tensorflow as tf
import json
import csv
import sys
import random
import numpy as np

train_columns = [
    "CPU",
    "Memory"
]

test_columns = [
    "word",
    "label"
]

categorical_columns = [
    "word"
]
continuous_columns = []

labels = []

def parse_json():
    with open(file) as f:
        data = json.load(f)

    train_file = csv.writer(open("train.csv", "wb+"))
    test_file = csv.writer(open("test.csv", "wb+"))
    unlabeled_file = csv.writer(open("unlabeled.csv", "wb+"))

    """
    Write headers
    train_file.writerow(train_columns)
    test_file.writerow(test_columns)
    """

    # Since the Tensorflow DNN (Deep Neural Network) wants the response variable in terms of numbers
    # Here we will change the labels into numbers (their respective indices)

    global labels
    labels = [value for color, value in data["labelNames"].iteritems()]
    print("Labels", labels)

    # Now let's subset the data by the individual cards and shuffle them
    card_data = data['cards']
    random.shuffle(card_data)

    index = 0

    for card in card_data:
        # Here we will split up the data into 2/3 training and 1/3 test
        ratio = 3

        if len(card['labels']) == 0:
            print(card['id'])
            #unlabeled_file.writerow([card['id'], tokenize_row_write(unlabeled_file, card['name'], card['desc'], "")])
            #unlabeled_file.writerow([card['id'], card['name'], ""])
            tokenize_row_write(unlabeled_file, card['id'], card['name'], card['desc'], "")
            continue

        write_to_file = None
        if index % ratio == 0:
            write_to_file = test_file
        else:
            write_to_file = train_file

        label = card['labels'][0]['name']
        tokenize_row_write(write_to_file, "", card['name'], card['desc'], labels.index(label))

        index += 1

def categorize(model_dir, model):
    word_hashed = tf.contrib.layers.sparse_column_with_hash_bucket("word", hash_bucket_size = 1000)
    # label_hashed = tf.contrib.layers.sparse_column_with_keys(column_name="label", keys=labels)

    # Now we will make our sets of wide and deep columns
    wide_columns = [
        word_hashed,
        # label_hashed,
        # tf.contrib.layers.crossed_column([word_hashed, label_hashed], hash_bucket_size=int(1e4))
    ]
    deep_columns = [
        tf.contrib.layers.embedding_column(word_hashed, dimension = 1),
    ]

    # These are the number of classes we are trying to predict
    num_classes = 10

    # Here we will build a Logistic Regression Model or a Deep Neural Network Classifier depending on need

    if model == "wide":
        m = tf.contrib.learn.LinearClassifier(model_dir=model_dir, feature_columns=wide_columns)
    elif model == "deep":
        m = tf.contrib.learn.DNNClassifier(model_dir=model_dir, feature_columns=deep_columns, hidden_units=[25, 10], n_classes=num_classes)
    elif model == "both":
        m = tf.contrib.learn.DNNLinearCombinedClassifier(
            model_dir=model_dir,
            linear_feature_columns=wide_columns,
            dnn_feature_columns=deep_columns,
            dnn_hidden_units=[25, 10],
            n_classes=num_classes)

    return m

def input_func(df):
    # Create a dictionary mapping from each continuous feature column name (k) to the values of that column stored in a constant Tensor.

    continuous_cols = {k: tf.constant(df[k].values)
                       for k in continuous_columns}

    # Creates a dictionary mapping from each categorical feature column name (k)
    # to the values of that column stored in a tf.SparseTensor.

    categorical_cols = {k: tf.SparseTensor(
        indices = [[i, 0] for i in range(df[k].size)],
        values = df[k].values,
        shape = [df[k].size, 1])
                        for k in categorical_columns}

    # Merge two dictionaries into one
    feature_cols = dict(continuous_cols)
    feature_cols.update(categorical_cols)

    # Converts the label column into a constant Tensor with dtype tf.int64
    label = tf.constant(df["label"].values)
    label = tf.to_int64(label)

    # Returns the feature columns and the label
    return feature_cols, label




def train_and_evaluate(model_call, card_name):
    # With this function, we will be parsing through Trello cards, getting the tokenized words of the title/description (while throwing out stop words)
    # For each of the tokenized words we will determine if existing labels apply
    # And then predict the current category based on previous categories
    # Now to train and evaluate model
    train_file = "train.csv"
    test_file = "test.csv"

    df_train = pd.read_csv(train_file, names = train_columns)
    df_test = pd.read_csv(test_file, names = test_columns)

    print(df_train.shape)
    print(df_test.shape)

    print(df_train.columns)
    print(df_test.columns)

    model_dir = tempfile.mkdtemp()

    m = categorize(model_dir, model_call)

    m.fit(input_fn = lambda: input_func(df_train), steps = 200)

    results = m.evaluate(input_fn = lambda: input_func(df_test), steps = 1)
    for key in sorted(results):
        print("{}: {}".format(key, results[key]))


    df_unlabeled = pd.read_csv("unlabeled.csv", names = ['id', 'word', 'label'])

    print(df_unlabeled.shape)
    print(df_unlabeled.columns)

    y = m.predict(input_fn = lambda: input_func(df_unlabeled))
    print ('Predictions: {}'.format(str(y)))


if __name__ == '__main__':
    parse_json()
    train_and_evaluate("both")
    #train_and_evaluate("deep")
    #train_and_evaluate("wide")
