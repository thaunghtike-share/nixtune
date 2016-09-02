# Copyright (C) 2016 Acksin <hey@acksin.com>
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

"""
We are looking at various metrics and computer usage:

1. Network Utilization
2. In/Out Network traffic
3. Stats
4. Memory usage
5. CPU usage
6. IO Usage
7. Primary Process

We will be using the Wide and Deep model provided by Tensorflow to
provide recommendations.
"""

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
    "Network Utilization",
    "Input Traffic",
    "Output Traffic",
    "Memory",
    "CPU",
    "IO Usage",
    "Primary Process"
]

test_columns = [
    "Network Utilization",
    "Input Traffic",
    "Output Traffic",
    "Memory",
    "CPU",
    "IO Usage",
    "Primary Process"
]

categorical_columns = [
    "Primary Process"
]

continuous_columns = [
    "Network Utilization",
    "Input Traffic",
    "Output Traffic",
    "Memory",
    "CPU",
    "IO Usage",
]

inst_types = []

file = "config.json"

azs = [
    "ap-northeast-1a",
    "ap-northeast-1c",
    "ap-northeast-2a",
    "ap-northeast-2c",
    "ap-southeast-1a",
    "ap-southeast-1b",
    "ap-southeast-2a",
    "ap-southeast-2b",
    "ap-southeast-2c",
    "eu-central-1a",
    "eu-central-1b",
    "eu-west-1a",
    "eu-west-1b",
    "eu-west-1c",
    "sa-east-1a",
    "sa-east-1c",
    "us-east-1a",
    "us-east-1b",
    "us-east-1d",
    "us-east-1e",
    "us-west-1a",
    "us-west-1c",
    "us-west-2a",
    "us-west-2b",
    "us-west-2c",
]

aws_inst_types = [
    "c1.medium",
    "c1.xlarge",
    "c3.2xlarge",
    "c3.4xlarge",
    "c3.8xlarge",
    "c3.large",
    "c3.xlarge",
    "c4.2xlarge",
    "c4.4xlarge",
    "c4.8xlarge",
    "c4.large",
    "c4.xlarge",
    "cc2.8xlarge",
    "cg1.4xlarge",
    "cr1.8xlarge",
    "d2.2xlarge",
    "d2.4xlarge",
    "d2.8xlarge",
    "d2.xlarge",
    "g2.2xlarge",
    "g2.8xlarge",
    "hi1.4xlarge",
    "i2.2xlarge",
    "i2.4xlarge",
    "i2.8xlarge",
    "i2.xlarge",
    "m1.large",
    "m1.medium",
    "m1.small",
    "m1.xlarge",
    "m2.2xlarge",
    "m2.4xlarge",
    "m2.xlarge",
    "m3.2xlarge",
    "m3.large",
    "m3.medium",
    "m3.xlarge",
    "m4.10xlarge",
    "m4.2xlarge",
    "m4.4xlarge",
    "m4.large",
    "m4.xlarge",
    "r3.2xlarge",
    "r3.4xlarge",
    "r3.8xlarge",
    "r3.large",
    "r3.xlarge",
    "t1.micro",
    "x1.32xlarge",
]

gcp_inst_types = [
    "n1-standard-1",
    "n1-standard-2",
    "n1-standard-4",
    "n1-standard-8",
    "n1-standard-16",
    "n1-standard-32",
    "n1-highmem-2",
    "n1-highmem-4",
    "n1-highmem-8",
    "n1-highmem-16",
    "n1-highmem-32",
    "n1-highcpu-2",
    "n1-highcpu-4",
    "n1-highcpu-8",
    "n1-highcpu-16",
    "n1-highcpu-32",
    "f1-micro",
    "g1-small",
]

azure_inst_types = [
    "A0",
    "A1",
    "A2",
    "A3",
    "A4",
    "A5",
    "A6",
    "A7",
    "A8",
    "A9",
    "A10",
    "A11",
    "D1",
    "D2",
    "D3",
    "D4",
    "D11",
    "D12",
    "D13",
    "D14",
    "D1v2",
    "D2v2",
    "D3v2",
    "D4v2",
    "D5v2",
    "D11v2",
    "D12v2",
    "D13v2",
    "D14v2",
    "D15v2",
    "F1",
    "F2",
    "F4",
    "F8",
    "F16",
    "G1",
    "G2",
    "G3",
    "G4",
    "G5"
]

instance_providers = [
    "AWS",
    "GCP",
    "Azure"
]

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

    # Since the Tensorflow DNN (Deep Neural Network) wants the
    # response variable in terms of numbers Here we will change the
    # instance types into numbers (their respective indices)

    global inst_types
    inst_types = [value for value in inst_type.iteritems()]
    print("Instance Types", inst_types)

    # Now let's get the data we need
    # First the memory data
    memory_data = data['System']['Memory']

    """
    What types are we concerned with? Free? Cached? Etc.
    """

    # Primary Processes
    process_data = data['Processes']

    """
    We looking at all of them or what??
    """

    # CPU Usage
    cpu_data = data["CPU"]

    # Data that has not been added yet
    # Placeholder names being used

    """
    net_util_data = data["Network Utilization"]
    in_traffic_data = data["Input Traffic"]
    out_traffic_data = data["Output Traffic"]
    io_usage_data = data["IO Usage"]
    """

    index = 0

    """
    Here's some pseudocode

    inst_type_provider = ...
    poss_inst_types = ...
    if inst_type_provider = "AWS":
        poss_inst_types = aws_inst_types
    if inst_type_provider = "GCP":
        poss_inst_types = gcp_inst_types
    if inst_type_provider = "Azure":
        poss_inst_types = azure_inst_types


    if memory_data < 1:
        ...
    if process_data == ...:
        ...
    if cpu_data < 5:
        ...
    """

def categorize(model_dir, model):
    # Categorical Columns
    process_hashed = tf.contrib.layers.sparse_column_with_hash_bucket(
        "process_data",
        hash_bucket_size = 1000
    )
    # Continuous Columns
    cpu_hashed = tf.contrib.layers.real_valued_column("cpu_data")
    memory_hashed = tf.contrib.layers.real_valued_column("memory_data")
    net_util_hashed = tf.contrib.layers.real_valued_column("net_util_data")
    in_traffic_hashed = tf.contrib.layers.real_valued_column("in_traffic_data")
    out_traffic_hashed = tf.contrib.layers.real_valued_column("out_traffic_data")
    io_usage_hashed = tf.contrib.layers.real_valued_column("io_usage_data")



    # Now we will make our sets of wide and deep columns
    # Need to add crossed columns

    wide_columns = [
        process_hashed,
    ]

    deep_columns = [
        tf.contrib.layers.embedding_column(
            process_hashed,
            dimension = 1
        ),
        cpu_hashed,
        memory_hashed,
        net_util_hashed,
        in_traffic_hashed,
        out_traffic_hashed,
        io_usage_hashed
    ]

    # These are the number of instance types we are trying to predict
    num_classes = 48

    # Here we will build a Logistic Regression Model or a Deep Neural
    # Network Classifier depending on need

    if model == "wide":
        m = tf.contrib.learn.LinearClassifier(
            model_dir = model_dir,
            feature_columns = wide_columns
        )
    elif model == "deep":
        m = tf.contrib.learn.DNNClassifier(
            model_dir = model_dir,
            feature_columns = deep_columns,
            hidden_units = [25, 10],
            n_classes = num_classes
        )
    elif model == "both":
        m = tf.contrib.learn.DNNLinearCombinedClassifier(
            model_dir = model_dir,
            linear_feature_columns = wide_columns,
            dnn_feature_columns = deep_columns,
            dnn_hidden_units = [25, 10],
            n_classes = num_classes)

    return m

def input_func(df):
    # Create a dictionary mapping from each continuous feature column
    # name (k) to the values of that column stored in a constant
    # Tensor.

    continuous_cols = {k: tf.constant(df[k].values)
                       for k in continuous_columns}

    # Creates a dictionary mapping from each categorical feature
    # column name (k) to the values of that column stored in a
    # tf.SparseTensor.

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
    # With this function, we will be parsing through Trello cards,
    # getting the tokenized words of the title/description (while
    # throwing out stop words) For each of the tokenized words we will
    # determine if existing labels apply And then predict the current
    # category based on previous categories Now to train and evaluate
    # model
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
    if len(sys.argv == 2):
        if sys.argv[1] == 'both':
            train_and_evaluate("both")
        if sys.argv[1] == 'deep':
            train_and_evaluate("deep")
        if sys.argv[1] == 'wide':
            train_and_evaluate("wide")
