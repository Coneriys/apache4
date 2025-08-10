#!/usr/bin/env bash
set -e

if [ -n "${VERSION}" ]; then
  echo "Deploying..."
else
  echo "Skipping deploy"
  exit 0
fi

git config --global user.email "${apache4ER_EMAIL}"
git config --global user.name "apache4er"

# load ssh key
eval "$(ssh-agent -s)"
chmod 600 ~/.ssh/apache4er_rsa
ssh-add ~/.ssh/apache4er_rsa

# update apache4-library-image repo (official Docker image)
echo "Updating apache4-library-imag repo..."
git clone git@github.com:apache4/apache4-library-image.git
cd apache4-library-image
./updatev2.sh "${VERSION}"
git add -A
echo "${VERSION}" | git commit --file -
echo "${VERSION}" | git tag -a "${VERSION}" --file -
git push -q --follow-tags -u origin master > /dev/null 2>&1

cd ..
rm -Rf apache4-library-image/

echo "Deployed"
